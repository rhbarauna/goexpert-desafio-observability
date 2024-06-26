package usecase

import (
	"context"

	"pos-graduacao/desafios/observabilidade/input/internal/entity"
	"pos-graduacao/desafios/observabilidade/input/internal/infra/forecast"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type ForecastOutputDTO struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

type GetForecast struct {
	forecastProvider forecast.ForecastProviderInterface
	tracer           trace.Tracer
}

func NewGetForecastUseCase(forecastProvider forecast.ForecastProviderInterface, tracer trace.Tracer) GetForecast {
	return GetForecast{forecastProvider, tracer}
}

func (uc *GetForecast) Execute(cep string, ctx context.Context) (ForecastOutputDTO, error) {
	ctx, span := uc.tracer.Start(ctx, "Normalizing Postal Code")
	span.SetAttributes(attribute.String("cep", cep))
	defer span.End()

	normalized, err := entity.NormalizePostalCode(cep)
	outputDTO := ForecastOutputDTO{}

	if err != nil {
		return outputDTO, err
	}

	forecast, err := uc.forecastProvider.GetForecast(normalized, ctx)

	if err != nil {
		return outputDTO, err
	}

	outputDTO.City = forecast.City
	outputDTO.TempC = forecast.TempC
	outputDTO.TempF = forecast.TempF
	outputDTO.TempK = forecast.TempK

	return outputDTO, nil
}
