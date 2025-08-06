package models

import (
	"encoding/json"
	"time"
)

type OutgoingEvent struct {
	ID          string          `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Topic       string          `gorm:"column:topic;not null"`
	Key         string          `gorm:"column:key"`
	Payload     json.RawMessage `gorm:"column:payload;type:jsonb;not null"`
	RetryCount  int             `gorm:"column:retry_count;default:0"`
	MaxRetries  int             `gorm:"column:max_retries;default:3"`
	CreatedAt   time.Time       `gorm:"column:created_at;default:now()"`
	IsPublished bool            `gorm:"column:is_published;default:false"`
}
