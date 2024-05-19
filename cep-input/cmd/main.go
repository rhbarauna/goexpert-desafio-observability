package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	shutdown := NewTracing("http://zipkin:9411/api/v2/spans", "cep-input")
	defer shutdown()

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	cepHandler := NewCepForecastHandler()
	router.MethodFunc(http.MethodGet, "/", cepHandler.Handle)

	log.Println("Iniciando o servidor web...")
	http.ListenAndServe(":8080", router)
}
