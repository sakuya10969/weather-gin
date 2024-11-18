package routes

import (
	"github.com/gin-gonic/gin"
	"api/handlers"
	"api/repositories"
	"api/services"
)

// SetupRouter: ルーティングをセットアップする関数
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// リポジトリ、サービス、ハンドラーのインスタンス作成
	weatherRepo := repositories.NewWeatherRepository()
	weatherService := services.NewWeatherService(weatherRepo)
	weatherHandler := handlers.NewWeatherHandler(weatherService)

	// ルートを定義
	r.GET("/weather", weatherHandler.GetWeather)

	return r
}