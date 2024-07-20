package model

import "github.com/google/uuid"

type Wholesale struct {
	ID   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name string

	ProductPrices []ProductPrice `gorm:"foreignKey:WholesaleID"`
}
