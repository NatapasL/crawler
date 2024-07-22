package productupdater

import (
	"github.com/google/uuid"
)

type ProductRepository interface {
	FindByWholesaleIDAndExternalID(wholesaleID uuid.UUID, externalID string) (*Product, error)
	Persist(Product) (uuid.UUID, error)
}
