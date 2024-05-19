package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	shutdown := NewTracing("http://zipkin:9411/api/v2/spans", "weather-api")
	defer shutdown()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	cepHandler := NewGetPlaceTemperaturesHandler()
	router.MethodFunc(http.MethodGet, "/", cepHandler.Handle)

	log.Println("Iniciando o servidor web...")
	http.ListenAndServe(":8080", router)
}
