package wholesale

import "github.com/google/uuid"

type Wholesale struct {
	id uuid.UUID
}

type ExistingWholesaleArgs struct {
	ID uuid.UUID
}

func ExistingWholesale(args ExistingWholesaleArgs) (*Wholesale, error) {
	ws := Wholesale{id: args.ID}
	err := ws.validate()
	if err != nil {
		return nil, err
	}

	return &ws, nil
}

func (Wholesale) validate() error {
	return nil
}

func (ws Wholesale) ID() uuid.UUID {
	return ws.id
}
