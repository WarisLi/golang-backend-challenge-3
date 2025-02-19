package main

import (
	"github.com/WarisLi/golang-backend-challenge-3/adapters"
	"github.com/WarisLi/golang-backend-challenge-3/core"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	beefRepo := adapters.NewAPIClient()
	beefService := core.NewBeefService(beefRepo)
	beefHandler := adapters.NewHttpBeefHandler(beefService)
	app.Get("/beef/summary", beefHandler.GetBeefs)

	app.Listen(":8080")
}
