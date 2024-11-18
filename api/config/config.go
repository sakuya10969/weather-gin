package config

import (
    "os"
)

var OpenWeatherAPIKey = os.Getenv("OPENWEATHER_API_KEY")