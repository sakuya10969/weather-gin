package handlers

import (
	"api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type WeatherHandler struct {
	weatherService *services.WeatherService
}

func NewWeatherHandler(s *services.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		weatherService: s,
	}
}

func (h *WeatherHandler) GetWeather(c *gin.Context) {
	city := c.Query("city")
	apikey := os.Getenv("OPENWEATHER_API_KEY")

	weather, err := h.weatherService.GetWeather(city, apikey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, weather)
}