package views

import (
	"os"

	// rice "github.com/GeertJohan/go.rice"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gofiber/template/html"
)

func HtmlEngine() *html.Engine {
	var engine *html.Engine
	debug := os.Getenv("STAGE_STATUS")
	if debug == "dev" {
		engine = html.New("./internal/adminApp/interfaces/views/templates", ".html")
		engine.Debug(true)
	} else {
		// load templates from rice-box
		engine = html.NewFileSystem(rice.MustFindBox("templates").HTTPBox(), ".html")
		engine.Debug(false)
	}
	return engine // *html.Engine

}
