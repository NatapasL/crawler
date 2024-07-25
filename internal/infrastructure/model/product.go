package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	SeriesID uuid.UUID `gorm:"type:uuid"`
	Name     string
	Volume   int

	CreatedAt time.Time
	UpdatedAt time.Time

	Prices []ProductPrice `gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return "product"
}
