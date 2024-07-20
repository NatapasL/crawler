package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductPrice struct {
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ProductID      uuid.UUID `gorm:"type:uuid"`
	WholesaleID    uuid.UUID `gorm:"type:uuid"`
	Price          float64
	EffectiveStart time.Time
	EffectiveEnd   time.Time
	SourceUrl      string
	Discounted     bool

	CreatedAt time.Time
	UpdatedAt time.Time

	Wholesale Wholesale
	Product   Product
}
