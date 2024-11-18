package repositories

import (
	"api/models"
	"encoding/json"
	"net/http"
)

const BASE_URL = "https://api.openweathermap.org/data/2.5/weather"

type WeatherRepository struct{}

func NewWeatherRepository() *WeatherRepository {
	return &WeatherRepository{}
}

func (r *WeatherRepository) GetWeather(city, apikey string) (*models.WeatherResponse, error) {

	url := BASE_URL + "?q=" + city + "&appid=" + apikey
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherResponse models.WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return nil, err
	}

	return &weatherResponse, nil
}