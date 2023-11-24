package controller

import "assignment-three/repository"

var (
	WeatherRepository repository.WeatherRepository
)

func init() {
	WeatherRepository = repository.InitWeatherReposository()
}
