package worker

import (
	"context"
	"log"
	"time"

	"github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
	"github.com/ngoctb13/seta-train/rest-service/internal/domains/team/repos"
	"github.com/ngoctb13/seta-train/shared-modules/config"
	"github.com/ngoctb13/seta-train/shared-modules/kafka"
)

type Worker struct {
	config            *config.AppConfig
	outgoingEventRepo repos.IOutgoingEventRepo
	producer          *kafka.Producer
}

func InitWorker(cfg *config.AppConfig, eventRepo repos.IOutgoingEventRepo, producer *kafka.Producer) *Worker {
	return &Worker{
		config:            cfg,
		outgoingEventRepo: eventRepo,
		producer:          producer,
	}
}

func (w *Worker) Start(ctx context.Context) {
	interval := time.Duration(w.config.Worker.Interval) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Worker shutting down")
			return
		case <-ticker.C:
			w.processBatch(ctx)
		}
	}
}

func (w *Worker) processBatch(ctx context.Context) {
	pe, err := w.outgoingEventRepo.GetPendingEvents(ctx)
	if err != nil {
		log.Printf("GetPendingEvents got fail: %v", err)
	}

	for _, e := range pe {
		if err := w.processEvent(ctx, e); err != nil {
			log.Printf("processEvent got fail: %v", err)
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
	}

	return w.outgoingEventRepo.MarkEventPublished(ctx, event.ID)
}
