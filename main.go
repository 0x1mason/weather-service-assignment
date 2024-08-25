package main

import (
	"api/pkg/handler"
	"log"
	"net/http"
)

func main() {
	rh := handler.RequestHandler{NoaaHost: "https://api.weather.gov"}
	router := http.NewServeMux()
	router.HandleFunc("GET /forecast/{coordinates}", rh.GetForecast)

	if err := http.ListenAndServe(":8888", router); err != nil {
		log.Fatal(err)
	}
}
