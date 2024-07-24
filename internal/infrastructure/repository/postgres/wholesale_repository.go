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

func (repository WholesaleRepository) Persist(ws wholesale.Wholesale) error {
	var savedProducts []model.Product
	var savedPrices []model.ProductPrice
	for _, product := range ws.Products() {
		savedProducts = append(savedProducts, model.Product{
			ID:       product.ID(),
			Name:     product.Name(),
			SeriesID: product.SeriesID(),
		})

		savedPrices = append(savedPrices, model.ProductPrice{
			ID:          product.PriceID(),
			Price:       product.Price(),
			Discounted:  product.Discounted(),
			ProductID:   product.ID(),
			ExternalID:  product.ExternalID(),
			WholesaleID: ws.ID(),
		})
	}

	err := repository.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(savedProducts).Error
		if err != nil {
			return err
		}

		err = tx.Save(savedPrices).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
