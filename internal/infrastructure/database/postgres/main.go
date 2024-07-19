package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func getDialector(config PostgresConnectionConfig) gorm.Dialector {
	dsn := buildConnectionString(config)
	return postgres.Open(dsn)
}

func GetConnection(config PostgresConnectionConfig) *gorm.DB {
	if db == nil {
		dialector := getDialector(config)
		gormDB, err := gorm.Open(dialector, &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
		if err != nil {
			log.Println(err)
			return nil
		}
		db = gormDB
	}

	return db
}

func buildConnectionString(config PostgresConnectionConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.DbName, config.Port, config.SslMode, config.TimeZone)
}
