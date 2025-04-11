package main

import (
	"fmt"

	"github.com/billy191/billy191-go-backend/config"
	database "github.com/billy191/billy191-go-backend/pkg/database/migration"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hello world")

	app := fiber.New()
	cfg := config.GetValue()
	db := database.NewDatabase(
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
		cfg.PostgresUsername,
		cfg.PostgresPassword,
	)
	if err := db.Connect(); err != nil {
		panic(err)
	}
	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("hello world 🌈")
		return c.JSON(fiber.Map{
			"message": "hello world",
			"Status":  "success",
		})
	})

	app.Listen(":3000")
}
