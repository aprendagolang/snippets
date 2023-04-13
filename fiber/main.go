package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/:name", HelloHandler)
	app.Post("/person", PersonHandler)

	app.Listen(":3000")
}

func HelloHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Hello, %s!", c.Params("name")),
	})
}

type Person struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func PersonHandler(c *fiber.Ctx) error {
	var p Person
	if err := c.BodyParser(&p); err != nil {
		return err
	}

	return c.SendString(fmt.Sprintf("Nome: %s - Idade: %d\n", p.Name, p.Age))
}
