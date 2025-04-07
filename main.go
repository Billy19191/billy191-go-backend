package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Test struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	fmt.Println("hello world")

	app := fiber.New(fiber.Config{
		Network: fiber.NetworkTCP,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("hello world 🌈")
		return c.JSON(Test{
			Message: "hello world",
			Status:  "success",
		})
	})

	app.Listen(":3000")
}
