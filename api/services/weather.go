package services

import (
	"api/models"
	"api/repositories"
)

type WeatherService struct {
	weatherRepo *repositories.WeatherRepository
}

func NewWeatherService(r *repositories.WeatherRepository) *WeatherService {
	return &WeatherService{
		weatherRepo: r,
	}
}

func (s *WeatherService) GetWeather(city, apikey string) (*models.WeatherResponse, error) {
	return s.weatherRepo.GetWeather(city, apikey)
}