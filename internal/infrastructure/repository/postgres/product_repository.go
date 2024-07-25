package postgres

import (
	"manga-crawler/internal/domain/catalogupdate/product"
	"manga-crawler/internal/infrastructure/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (repository ProductRepository) Persist(p product.Product) error {
	savedProduct := model.Product{
		ID:       p.ProductID(),
		Name:     p.Name(),
		SeriesID: p.SeriesID(),
	}

	savedPrice := model.ProductPrice{
		ID:          p.PriceID(),
		Price:       p.Price(),
		Discounted:  p.Discounted(),
		ProductID:   p.ProductID(),
		ExternalID:  p.ExternalID(),
		WholesaleID: p.WholesaleID(),
	}

	err := repository.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(savedProduct).Error
		if err != nil {
			return err
		}

		err = tx.Save(savedPrice).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
