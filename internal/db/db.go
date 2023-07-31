package db

import (
	"log"

	"github.com/AlexS778/fqdnIPLookup/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBContext(connStr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.FQDN{}, &models.IP{})
	return db
}
