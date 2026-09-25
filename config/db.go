package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB adalah variabel global untuk mengakses instance GORM di package lain
var DB *gorm.DB

// ConnectDatabase menginisialisasi koneksi MySQL menggunakan GORM
func ConnectDatabase() {
	_ = godotenv.Load()

	dbUser := getEnv("DB_USER", "my_user")
	dbPass := getEnv("DB_PASSWORD", "my_password")
	dbHost := getEnv("DB_HOST", "127.0.0.1") // Gunakan mysql saat berjalan di Docker melalui DB_HOST
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "my_db")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke MySQL Database: %v", err)
	}

	log.Println("Berhasil terhubung ke database MySQL!")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
