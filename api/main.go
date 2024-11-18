package main

import (
	"api/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".envファイルの読み込み失敗")
	}

	r := routes.SetupRouter()
	r.Run(":8080")
}