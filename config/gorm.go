package config

import (
	"assignment-three/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

func ConnectGorm() *gorm.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	CheckError(err)

	// check db
	// err = db.Ping()
	CheckError(err)

	fmt.Println("Connected Gorm Sql!")

	db.AutoMigrate(&model.Weather{})

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
