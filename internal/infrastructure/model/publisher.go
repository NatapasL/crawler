package model

import (
	"time"

	"github.com/google/uuid"
)

type Publisher struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time

	Series []Series `gorm:"foreignKey:PublisherID"`
}

func (Publisher) TableName() string {
	return "publisher"
}
