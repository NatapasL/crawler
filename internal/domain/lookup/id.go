package lookup

import "github.com/google/uuid"

type ID string

func (id ID) UUID() uuid.UUID {
	return uuid.Must(uuid.Parse(string(id)))
}
