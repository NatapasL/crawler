package postgres

import (
	"errors"
	"manga-crawler/internal/domain/catalogupdate/wholesale"
	"manga-crawler/internal/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WholesaleRepository struct {
	db *gorm.DB
}

func NewWholesaleRepository(db *gorm.DB) *WholesaleRepository {
	return &WholesaleRepository{db}
}

func (repository WholesaleRepository) FindById(id uuid.UUID) (*wholesale.Wholesale, error) {
	var ws model.Wholesale
	err := repository.db.Where("id = ?", id).First(&ws).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return wholesale.ExistingWholesale(wholesale.ExistingWholesaleArgs{ID: ws.ID})
}
