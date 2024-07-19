package postgres

import (
	"manga-crawler/internal/infrastructure/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Series{},
		&model.SeriesName{},
		&model.Genre{},
		&model.SeriesGenre{},
		&model.Person{},
		&model.PersonName{},
		&model.Publisher{},
		&model.SeriesNameMatcher{},
	)
}
