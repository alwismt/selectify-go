package main

import (
	"fmt"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.NewFileSystem(rice.MustFindBox("../../web/templates/customer").HTTPBox(), ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:   rice.MustFindBox("../../web/static/customer").HTTPBox(),
		MaxAge: 3600,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	// app.Use("/controlpanel/assets", filesystem.New(filesystem.Config{
	// 	Root:   rice.MustFindBox("../../internal/infrastructure/web/admin_statics").HTTPBox(),
	// 	MaxAge: 3600,
	// }))

	app.Use("/controlpanel", func(c *fiber.Ctx) error {
		// get the url from the request
		url := c.OriginalURL()
		return proxy.Do(c, "http://localhost:8081"+url)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Received request at /")
		return c.Render("index", nil)
	})

	app.Listen(":8089")
}
