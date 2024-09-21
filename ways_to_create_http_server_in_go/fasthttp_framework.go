package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/valyala/fasthttp"
)

type baseResponse struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type requestHandler struct {}

func (h *requestHandler) handle404Request(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusNotFound)
	res, _ := json.Marshal(
		baseResponse{
			Message: "not found",
			Data: map[string]interface{}{
				"detail": "request resource not found",
			},
		},
	)
	ctx.SetBody(res)

	log.Default().Println("request resource not found")
}

func (h *requestHandler) indexHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(http.StatusOK)
	res, _ := json.Marshal(
		baseResponse{
			Message: "http server is running with fasthttp framework",
			Data:    map[string]interface{}{"created_by": "ARTM2000"},
		},
	)
	ctx.SetBody(res)
}

func (h *requestHandler) HandleRequests(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		h.indexHandler(ctx)

	default:
		h.handle404Request(ctx)
	}
}

func main() {
	h := requestHandler{}

	defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered. error:\n", r)
        }
    }()

	log.Default().Println("server listen on port 3000!")
	err := fasthttp.ListenAndServe(":3000", h.HandleRequests)
	if err != nil {
		log.Fatalf("server failed to start: %s", err.Error())
	}
}

