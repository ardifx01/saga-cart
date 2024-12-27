package db

import (
	"log"
	"os"
	"product_service_saga/internal/domain"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectPostgre() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=pepega90 dbname=db_saga_product port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second, // Log SQL queries slower than this
				LogLevel:      logger.Info, // Log level: Silent, Error, Warn, Info
				Colorful:      true,
			},
		),
	})

	db.AutoMigrate(&domain.Product{})

	if err != nil {
		log.Printf("error establish connection to postgresSQL: %v", err.Error())
		return nil, err
	}

	log.Println("Connected to postgreSQL via GORM")
	return db, nil
}
