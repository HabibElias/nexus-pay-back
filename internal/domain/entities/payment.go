package entities

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID        string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
