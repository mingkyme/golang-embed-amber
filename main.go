package main

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/amber"
)

//go:embed views/*
var views embed.FS

func main(){
	
	engine := amber.NewFileSystem(http.FS(views), ".amber")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	engine.Debug(true)
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("views/index",fiber.Map{})
    })

    app.Listen(":8080")

}