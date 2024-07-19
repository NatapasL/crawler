package model

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Names []PersonName `gorm:"foreignKey:PersonID"`
}
