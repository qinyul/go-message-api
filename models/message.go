package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID             uuid.UUID `json:"id" gorm:"primaryKey"`
	ConversationID uuid.UUID `json:"conversation_id"`
	SenderId       uuid.UUID `json:"sender_id"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	Body           string    `json:"body,omitempty"`
	IsSoftDeleted  bool      `json:"is_soft_deleted,omitempty"`
}
