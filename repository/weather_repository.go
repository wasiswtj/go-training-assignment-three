package repository

import "assignment-three/model"

type WeatherRepository interface {
	Update(model.Weather) (model.Weather, error)
}

type weatherRepository struct{}

func InitWeatherReposository() *weatherRepository {
	return &weatherRepository{}
}

func (weatherRepository) Update(weather model.Weather) (model.Weather, error) {
	err := DB.Create(&weather).Error
	if err != nil {
		return model.Weather{}, err
	}

	return weather, nil
}
