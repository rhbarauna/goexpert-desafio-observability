package handler

import (
	"encoding/json"
	"net/http"

	"pos-graduacao/desafios/observabilidade/input/internal/entity"
	"pos-graduacao/desafios/observabilidade/input/internal/infra/forecast"
	"pos-graduacao/desafios/observabilidade/input/internal/usecase"
)

type GetCepForecastHandler struct {
	uc usecase.GetForecast
}

func NewCepForecastHandler(uc usecase.GetForecast) GetCepForecastHandler {
	return GetCepForecastHandler{
		uc: uc,
	}
}

func (h *GetCepForecastHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	output, err := h.uc.Execute(cep, r.Context())

	if err != nil {
		if err == entity.ErrInvalidInputType || err == entity.ErrInvalidInputMinLength {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		if err == forecast.ErrZipCodeNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	err = json.NewEncoder(w).Encode(output)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
