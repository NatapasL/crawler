package model

import (
	"github.com/google/uuid"
)

type PersonName struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	PersonID   uuid.UUID `gorm:"type:uuid"`
	GivenName  string
	FamilyName string
	Language   string
}

func (PersonName) TableName() string {
	return "person_name"
}
