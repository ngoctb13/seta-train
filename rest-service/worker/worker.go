package worker

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
	"github.com/ngoctb13/seta-train/rest-service/internal/domains/team/repos"
	"github.com/ngoctb13/seta-train/shared-modules/config"
	"github.com/ngoctb13/seta-train/shared-modules/infra/kafka"
)

type Worker struct {
	config            *config.AppConfig
	outgoingEventRepo repos.IOutgoingEventRepo
	producer          *kafka.Producer
	jobChan           chan *models.OutgoingEvent
	wg                sync.WaitGroup
}

func InitWorker(cfg *config.AppConfig, eventRepo repos.IOutgoingEventRepo, producer *kafka.Producer) *Worker {
	return &Worker{
		config:            cfg,
		outgoingEventRepo: eventRepo,
		producer:          producer,
		jobChan:           make(chan *models.OutgoingEvent, cfg.Worker.BatchSize*2),
	}
}

func (w *Worker) Start(ctx context.Context) {
	// Start worker
	for i := 0; i < w.config.Worker.Concurrency; i++ {
		go w.startWorker(ctx, i)
	}

	interval := time.Duration(w.config.Worker.Interval) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Worker shutting down")
			close(w.jobChan)
			w.wg.Wait()
			return
		case <-ticker.C:
			w.dispatchEvents(ctx)
		}
	}
}

func (w *Worker) startWorker(ctx context.Context, workerID int) {
	log.Printf("[Worker-%d] Started", workerID)
	for event := range w.jobChan {
		w.wg.Add(1)

		if err := w.processEvent(ctx, event); err != nil {
			log.Printf("[Worker-%d] processEvent failed for event %s: %v", workerID, event.ID, err)
		}

		w.wg.Done()
	}
	log.Printf("[Worker-%d] Exited", workerID)
}

func (w *Worker) dispatchEvents(ctx context.Context) {
	events, err := w.outgoingEventRepo.GetPendingEvents(ctx, w.config.Worker.BatchSize)
	if err != nil {
		log.Printf("GetPendingEvents failed: %v", err)
		return
	}

	for _, event := range events {
		select {
		case w.jobChan <- event:
		case <-ctx.Done():
			return
		}
	}
}

func (w *Worker) processEvent(ctx context.Context, event *models.OutgoingEvent) error {
	// publish to kafka
	_, _, err := w.producer.SendMessage(ctx, event.Topic, event.Payload, kafka.ProducerMessageOption{
		Key: event.Key,
	})
	if err != nil {
		// increase retry count
		if retryErr := w.outgoingEventRepo.IncrementRetryCount(ctx, event.ID); retryErr != nil {
			log.Printf("IncrementRetryCount got fail: %v", retryErr)
		}
		return err
	}

	return w.outgoingEventRepo.MarkEventPublished(ctx, event.ID)
}
