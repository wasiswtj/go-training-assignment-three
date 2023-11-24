package model

import "gorm.io/gorm"

type Weather struct {
	gorm.Model
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
