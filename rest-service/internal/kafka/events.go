package kafka

import (
	"encoding/json"
	"time"
)

type TeamCreatedEvent struct {
	TeamID    string    `json:"team_id"`
	TeamName  string    `json:"team_name"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (e *TeamCreatedEvent) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}
