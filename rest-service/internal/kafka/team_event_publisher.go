package kafka

import (
	"context"
	"time"

	"github.com/ngoctb13/seta-train/rest-service/internal/domains/team/usecases"
	sharedKafka "github.com/ngoctb13/seta-train/shared-modules/kafka"
)

type TeamEventPublisherImpl struct {
	Producer *sharedKafka.Producer
}

func NewTeamEventPublisher(producer *sharedKafka.Producer) usecases.TeamEventPublisher {
	return &TeamEventPublisherImpl{
		Producer: producer,
	}
}

func (p *TeamEventPublisherImpl) TeamCreated(teamID, teamName, userID string) error {
	event := &TeamCreatedEvent{
		TeamID:    teamID,
		TeamName:  teamName,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	data, err := event.ToJSON()
	if err != nil {
		return err
	}

	_, _, err = p.Producer.SendMessage(context.Background(), "rest-service.team", data, sharedKafka.ProducerMessageOption{
		Key: teamID,
	})

	return err
}
