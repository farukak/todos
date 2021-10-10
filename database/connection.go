package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() {
	var err error

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSL_MODE")
	timeZone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timeZone=%s",
		host, user, password, dbName, port, sslMode, timeZone)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("db connection was unsuccessful.")
	}

	//postgresql, err := db.DB()
	//postgresql.SetMaxIdleConns(10)
	//postgresql.SetMaxOpenConns(10)
	//postgresql.SetConnMaxLifetime(time.Hour)

	//Database Migration
	//err = db.AutoMigrate()

	log.Println("db connection was successful.")
}
