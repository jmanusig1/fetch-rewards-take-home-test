package main 

import (
	"log"
	"golang-fetch-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) { 
    app.Get("/get-balances", handlers.GetBalances)
    app.Post("/add-transaction", handlers.AddTransaction )
    app.Post("/spend-points", handlers.SpendPoints)
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Fetch Points Api")
    })
}

func main() { 
    app := fiber.New()
    setupRoutes(app)
    log.Fatal(app.Listen(":5000"))
}