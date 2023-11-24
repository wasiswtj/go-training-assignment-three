package controller

import (
	"assignment-three/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateData(c echo.Context) error {
	var weather model.Weather
	if err := c.Bind(&weather); err != nil {
		return err
	}

	weather, err := WeatherRepository.Update(weather)
	if err != nil {
		return err
	}

	log.Println("Water", weather.Water, MapWindStatus(weather.Water))
	log.Println("Wind", weather.Wind, MapWindStatus(weather.Wind))

	return c.JSON(http.StatusOK, map[string]string{
		"status": "success",
	})
}

func MapWindStatus(wind int) string {
	if wind <= 6 {
		return "aman"
	} else if wind >= 7 && wind <= 15 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func MapWaterStstus(water int) string {
	if water <= 5 {
		return "aman"
	} else if water >= 6 && water <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}
