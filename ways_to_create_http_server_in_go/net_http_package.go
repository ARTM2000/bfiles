package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type baseResponse struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func handle404Request(w http.ResponseWriter) {
	// set `Content-Type` header
	w.Header().Add("Content-Type", "application/json")

	// set response status code
	w.WriteHeader(http.StatusNotFound)

	// encode response struct to json
	res, _ := json.Marshal(
		baseResponse{
			Message: "not found",
			Data: map[string]interface{}{
				"detail": "requested resource not found",
			},
		},
	)

	fmt.Fprint(w, string(res))

	log.Default().Println("request resource not found")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet || r.RequestURI != "/" {
		handle404Request(w)
		return
	}

	// encode response struct to json
	res, _ := json.Marshal(
		baseResponse{
			Message: "http server is running with native package",
			Data:    map[string]interface{}{"created_by": "ARTM2000"},
		},
	)

	// set `Content-Type` header
	w.Header().Add("Content-Type", "application/json")

	// set response status code
	w.WriteHeader(http.StatusOK)

	// write the response
	fmt.Fprint(w, string(res))
}

func main() {
	http.HandleFunc("/", indexHandler)

	defer func() {
        if r := recover(); r != nil {
            fmt.Println("panic recovered. error:\n", r)
        }
    }()

	log.Default().Println("server listen on port 3000!")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("server failed to start: %s", err.Error())
	}
}
