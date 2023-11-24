package repository

import (
	"assignment-three/config"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	DB = config.ConnectGorm()
}

func GetDB() *gorm.DB {
	return GetDB()
}
