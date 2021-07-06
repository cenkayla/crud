package model

import (
	"encoding/json"
	"time"
)

type Student struct {
	ID        json.Number `gorm:"primary_key" json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	Age       int         `json:"age,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
