package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
)

type IOutgoingEventRepo interface {
	CreateOutgoingEvent(ctx context.Context, event *models.OutgoingEvent) error
	GetPendingEvents(ctx context.Context) ([]*models.OutgoingEvent, error)
	MarkEventPublished(ctx context.Context, eventID string) error
	IncrementRetryCount(ctx context.Context, eventID string) error
}
