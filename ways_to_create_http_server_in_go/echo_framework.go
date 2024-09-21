package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type baseResponse struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func handle404Request(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusNotFound,
		baseResponse{
			Message: "not found",
			Data: map[string]interface{}{
				"detail": "request resource not found",
			},
		},
	)
}

func indexHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK,
		baseResponse{
			Message: "http server is running with echo framework",
			Data:    map[string]interface{}{"created_by": "ARTM2000"},
		},
	)
}

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", indexHandler)
	app.RouteNotFound("/*", handle404Request)

	log.Default().Println("server listen on port 3000!")
	app.Logger.Fatal(app.Start(":3000"))
}
