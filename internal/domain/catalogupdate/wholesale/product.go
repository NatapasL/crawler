package wholesale

import "github.com/google/uuid"

type Product interface {
	SetWholesaleID(uuid.UUID)
}
