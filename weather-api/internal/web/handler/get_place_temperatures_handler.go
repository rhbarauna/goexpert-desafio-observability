package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type GetPlaceTemperaturesHandler struct {
	uc usecase.GetPlaceForecast
}

func NewGetPlaceTemperaturesHandler(uc usecase.GetPlaceForecast) GetPlaceTemperaturesHandler {
	return GetPlaceTemperaturesHandler{
		uc: uc,
	}
}

func (h *GetPlaceTemperaturesHandler) Handle(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := otel.GetTextMapPropagator().Extract(r.Context(), carrier)

	cep := r.URL.Query().Get("cep")
	output, err := h.uc.Execute(cep, ctx)

	if err != nil {
		if err == usecase.ErrInvalidInput {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		if err == usecase.ErrPostalCodeNotFound || err == usecase.ErrWeatherNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err == usecase.ErrPostalCodeNotFound || err == usecase.ErrWeatherNotFound {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
