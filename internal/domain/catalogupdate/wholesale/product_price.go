package wholesale

import "github.com/google/uuid"

type ProductPrice interface {
	ID() uuid.UUID
	Price() float64
	Discounted() bool
}
