package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var db *gorm.DB

var dbPort = getDontEnv("DB_PORT")
var dbUser = getDontEnv("DB_USER")
var dbPass = getDontEnv("DB_PASSWORD")
var dbHost = getDontEnv("DB_HOST")
var dbTable = getDontEnv("DB_TABLE")

func StartDB() {
	dns := dbUser + ":" + dbPass + "@tcp" + "(" + dbHost + ":" + dbPort + ")/" + dbTable + "?" + "parseTime=true&loc=Local"

	database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB {
	return db
}

func getDontEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
