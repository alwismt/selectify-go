package config

import (
	"os"
	"time"

	"github.com/alwismt/selectify/internal/infrastructure/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

func Config(app *fiber.App) {

	// setup compress
	app.Use(compress.New(
		compress.Config{
			Level: compress.LevelBestSpeed, // set compression level
		},
	))

	ck_http_only := os.Getenv("COOKIE_HTTP_ONLY") // get cookie http only from env
	var ck_http_only_bool bool                    // set cookie http only bool
	if ck_http_only == "true" {
		ck_http_only_bool = true // set cookie http only bool to true
	} else {
		ck_http_only_bool = false // set cookie http only bool to false
	}

	// setup csrf
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "cookie:csrf_token",
		CookieName:     "csrf_token",
		CookieSameSite: "Strict",
		CookieHTTPOnly: ck_http_only_bool,
		Storage:        session.NewStore(),
		CookieSecure:   false,
		Expiration:     10 * time.Minute,
		KeyGenerator:   utils.UUIDv4,
	}))

}
