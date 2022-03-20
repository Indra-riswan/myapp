package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Erorr Load Env")
	}

	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dbhost, dbport, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erorr Connect DB")
	}

	return db

}

func CloseConnectionDB(db *gorm.DB) {
	Closedb, err := db.DB()
	if err != nil {
		log.Fatal("Failed Close Connection")
	}
	Closedb.Close()
}
