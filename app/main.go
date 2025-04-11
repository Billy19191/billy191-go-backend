package main

import (
	"fmt"

	// database "pkg/database/migration"

	"github.com/billy191/billy191-go-backend/config"
	database "github.com/billy191/billy191-go-backend/pkg/database/migration"
	"github.com/gofiber/fiber/v2"
)

type Test struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	fmt.Println("hello world")

	app := fiber.New()
	cfg := config.GetValue()
	err := database.NewDatabase(
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
		cfg.PostgresUsername,
		cfg.PostgresPassword,
	).Connect()
	if err != nil {
		panic(err)
	}
	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("hello world 🌈")
		return c.JSON(Test{
			Message: "hello world",
			Status:  "success",
		})
	})

	app.Listen(":3000")
}
