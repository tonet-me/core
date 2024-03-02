package entity

import (
	"time"
)

type Visit struct {
	ID        string            `bson:"_id,omitempty" json:"id"`
	CardID    string            `bson:"card_id" json:"card_id"`
	UserAgent map[string]string `bson:"user_agent" json:"user_agent"`
	CreatedAt time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time         `bson:"updated_at" json:"updated_at"`
}
