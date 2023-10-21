package db

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Unidade/fullcycle-desafio/domain/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}


}

func ConnectDB() *gorm.DB {
	var db *gorm.DB
	var err error

	dsn := os.Getenv("dsn")

	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database")
	}

	if os.Getenv("debug") == "true" {
		db.Logger.LogMode(logger.Info)
	}

	db.AutoMigrate(&model.Product{})
	

	return db
}
