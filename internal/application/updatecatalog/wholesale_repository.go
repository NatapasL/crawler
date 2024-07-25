package updatecatalog

import (
	"manga-crawler/internal/domain/catalogupdate/wholesale"

	"github.com/google/uuid"
)

type wholesaleRepository interface {
	FindById(uuid.UUID) (*wholesale.Wholesale, error)
}
