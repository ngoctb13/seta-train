package repos

import (
	"context"

	"github.com/ngoctb13/seta-train/rest-service/internal/domain/models"
	"gorm.io/gorm"
)

type outgoingEventSQLRepo struct {
	db *gorm.DB
}

func NewOutgoingEventSQLRepo(db *gorm.DB) *outgoingEventSQLRepo {
	return &outgoingEventSQLRepo{db: db}
}

func (o *outgoingEventSQLRepo) CreateOutgoingEvent(ctx context.Context, event *models.OutgoingEvent) error {
	return o.db.WithContext(ctx).Create(event).Error
}
func (o *outgoingEventSQLRepo) GetPendingEvents(ctx context.Context, limit int) ([]*models.OutgoingEvent, error) {
	var events []*models.OutgoingEvent
	err := o.db.WithContext(ctx).
		Where("is_published = false AND retry_count < max_retries").
		Order("created_at ASC").
		Limit(limit).
		Find(&events).Error

	return events, err
}

func (o *outgoingEventSQLRepo) MarkEventPublished(ctx context.Context, eventID string) error {
	return o.db.WithContext(ctx).
		Model(&models.OutgoingEvent{}).
		Where("id = ?", eventID).
		Update("is_published", true).Error
}

func (o *outgoingEventSQLRepo) IncrementRetryCount(ctx context.Context, eventID string) error {
	return o.db.WithContext(ctx).
		Model(&models.OutgoingEvent{}).
		Where("id = ?", eventID).
		Update("retry_count", gorm.Expr("retry_count + 1")).Error
}
