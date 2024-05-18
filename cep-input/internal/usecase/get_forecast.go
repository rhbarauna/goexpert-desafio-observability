package usecase

import (
	"errors"

	"pos-graduacao/desafios/observabilidade/input/internal/entity"
	"pos-graduacao/desafios/observabilidade/input/internal/infra/forecast"
)

type ForecastOutputDTO struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

type GetForecast struct {
	forecastProvider forecast.ForecastProviderInterface
}

func NewGetForecastUseCase(forecastProvider forecast.ForecastProviderInterface) GetForecast {
	return GetForecast{forecastProvider}
}

func (uc *GetForecast) Execute(cep string) (ForecastOutputDTO, error) {
	normalized, err := entity.NormalizePostalCode(cep)
	outputDTO := ForecastOutputDTO{}

	if err != nil {
		return outputDTO, err
	}

	forecast, err := uc.forecastProvider.GetForecast(normalized)

	if err != nil {
		return outputDTO, errors.New("Erro ao obter previsão do tempo.")
	}

	outputDTO.City = forecast.City
	outputDTO.TempC = forecast.TempC
	outputDTO.TempF = forecast.TempF
	outputDTO.TempK = forecast.TempK

	return outputDTO, nil
}
