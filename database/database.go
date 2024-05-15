package database

import (
	"log"

	"go-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    dsn := "user=postgres dbname=karyawan_v2 sslmode=disable password=postgres"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    db.AutoMigrate(&models.Karyawan{})
    DB = db
}
