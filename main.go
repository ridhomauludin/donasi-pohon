package main

import (
	"donasiPohon/config"
	"donasiPohon/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	e := echo.New()
	config.InitDB()
	config.InitGemini()

	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
