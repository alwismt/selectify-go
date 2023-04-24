package web

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/alwismt/selectify/internal/adminApp/config"
	"github.com/alwismt/selectify/internal/adminApp/interfaces/routes/api"
	"github.com/alwismt/selectify/internal/adminApp/interfaces/routes/web"
	"github.com/alwismt/selectify/internal/adminApp/interfaces/views"

	"github.com/alwismt/selectify/internal/infrastructure/messagebroker"
	"github.com/alwismt/selectify/internal/infrastructure/messagebroker/queue"
	mongodb "github.com/alwismt/selectify/internal/infrastructure/persistence/mongoDB"
	"github.com/alwismt/selectify/internal/infrastructure/persistence/redis"
	sqldb "github.com/alwismt/selectify/internal/infrastructure/persistence/sqlDB"
	"github.com/alwismt/selectify/internal/infrastructure/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func AdminServerStart() {
	stage := os.Getenv("STAGE_STATUS")
	var debug bool = false
	if stage == "dev" {
		debug = true
	} else {
		debug = false
	}

	// create new fiber app
	app := fiber.New(fiber.Config{
		Views:                 views.HtmlEngine(),
		PassLocalsToViews:     true,
		BodyLimit:             15 * 1024 * 1024,
		ViewsLayout:           "layouts/main",
		AppName:               "Selectify",
		DisableStartupMessage: !debug,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.HandleError(c, err)
		},
	})
	// setup config
	config.Config(app)

	// setup static files
	if debug {
		app.Use("/controlpanel/assets", filesystem.New(filesystem.Config{
			Root: http.Dir("./internal/infrastructure/web/admin_statics"),
		}))
	} else {
		app.Use("/controlpanel/assets", filesystem.New(filesystem.Config{
			Root:   rice.MustFindBox("admin_statics").HTTPBox(),
			MaxAge: 3600,
		}))
	}

	// initialize database connections
	sqldb.ConnectToDBwithPool()
	mongodb.ConnectToMongoDB()
	redis.CoonectToRedis()

	rmq := make(chan bool)
	go func() {
		if err := queue.RabbitMQConnection(); err != nil {
			panic(err)
		}
		rmq <- true
	}()
	// Wait for RabbitMQ connection to be established
	<-rmq

	// setup routes
	web.WebRoutes(app)
	api.APIRoutes(app)

	// set RabbitMQ
	// if err := queue.RabbitMQConnection(); err != nil {
	// 	panic(err)
	// }
	// Start the consumer listener
	go messagebroker.ConsumerListner()

	// numGoroutines := runtime.NumGoroutine()
	// fmt.Println("numGoroutines", numGoroutines)

	url, _ := utils.ConnectionURLBuilder("fiber")
	// listen
	if debug {
		err := app.Listen(url)
		if err != nil {
			panic(err)
		}
	} else {
		startServerWithGracefulShutdown(app, url)
	}
}

func startServerWithGracefulShutdown(a *fiber.App, url string) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.ShutdownWithTimeout(5 * time.Second); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(url); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
