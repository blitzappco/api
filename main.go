package main

import (
	"api/accounts"
	"api/db"
	"api/env"
	"api/tickets"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
  app := fiber.New()

  app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
  }))

  db.InitDB()
  db.InitCache(env.RedisOptions)

  app.Get("/test", func (c *fiber.Ctx) error {
    return c.SendString("it's working")
  })

  accounts.Routes(app)
  tickets.Routes(app)

  app.Listen(":9999")
}