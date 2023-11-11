package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	catfactsurl := os.Getenv("CAT_FACT_API_URL")

	getCatFactsHandler := NewCatFactService(catfactsurl).GetCatFact

	app := fiber.New()
	app.Use("/", getCatFactsHandler)
	app.Listen(":" + port)
}
