package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

const (
	FIBER_TIMEOUT = 20
)

var App *fiber.App
var API fiber.Router

func init() {
	App = fiber.New(fiber.Config{
		ReadTimeout:  FIBER_TIMEOUT * time.Second,
		WriteTimeout: FIBER_TIMEOUT * time.Second,
		IdleTimeout:  FIBER_TIMEOUT * time.Second,
		BodyLimit:    100 * 1024 * 1024,
	})
	App.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, X-API-KEY",
	}))
	App.Use(etag.New())
	App.Use(requestid.New())
	App.Use(logger.New())
	App.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	API = App.Group("/api")
	users_routes()
	class_routes()
	faculty_routes()
	Moderator_routes()
	batch_routes()
	examcategory_routes()
	youtube_routes()
}
