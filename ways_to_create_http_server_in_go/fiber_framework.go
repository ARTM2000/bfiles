package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type baseResponse struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func handle404Request(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusNotFound).JSON(
		baseResponse{
			Message: "not found",
			Data: map[string]interface{}{
				"detail": "request resource not found",
			},
		},
	)
}

func indexHandler(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(
		baseResponse{
			Message: "http server is running with fiber framework",
			Data:    map[string]interface{}{"created_by": "ARTM2000"},
		},
	)
}

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(logger.New())

	app.Get("/", indexHandler)

	// not found routes handler should add at the end to catch any unhandled routes
	app.Use(handle404Request)

	log.Default().Println("server listen on port 3000!")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("server failed to start: %s", err.Error())
	}
}
