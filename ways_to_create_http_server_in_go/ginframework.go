package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type baseResponse struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func handle404Request(ctx *gin.Context) {
	log.Default().Println("request resource not found")
	ctx.JSON(
		http.StatusNotFound,
		baseResponse{
			Message: "not found",
			Data: map[string]interface{}{
				"detail": "request resource not found",
			},
		},
	)
}

func indexHandler(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		baseResponse{
			Message: "http server is running with gin framework",
			Data:    map[string]interface{}{"created_by": "ARTM2000"},
		},
	)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/", indexHandler)
	router.NoRoute(handle404Request)

	log.Default().Println("server listen on port 3000!")
	router.Run(":3000")
}
