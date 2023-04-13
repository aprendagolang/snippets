package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

//go:embed views/*
var views embed.FS

func main() {
	// Initialize standard Go html template engine
	engine := html.NewFileSystem(http.FS(views), ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("views/index", fiber.Map{
			"Nome": "Tiago Temporin",
		})
	})

	log.Fatal(app.Listen(":3000"))
}

func DeferUm() {
	fmt.Println(1)
}
func DeferDois() {
	fmt.Println(2)
}
func DeferTres() {
	fmt.Println(3)
}

func ExemploDefer() {

	defer DeferUm()
	defer DeferDois()
	defer DeferTres()
}
