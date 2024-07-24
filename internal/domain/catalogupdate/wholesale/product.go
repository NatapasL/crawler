package wholesale

import "github.com/google/uuid"

type Product interface {
	SetWholesaleID(uuid.UUID)
	ID() uuid.UUID
	Name() string
	SeriesID() uuid.UUID
	ExternalID() string
	Price() float64
	Discounted() bool
	PriceID() uuid.UUID
}
