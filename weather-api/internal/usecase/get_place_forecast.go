package usecase

import (
	"context"
	"errors"

	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/entity"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/place"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/weather"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var (
	ErrPostalCodeNotFound = errors.New("zipcode not found")
	ErrWeatherNotFound    = errors.New("weather not found for that zipcode")
	ErrInvalidInput       = errors.New("invalid zipcode")
)

type PlaceForecastOutputDTO struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

type GetPlaceForecast struct {
	placeProvider   place.PlaceProviderInterface
	weatherProvider weather.WeatherProviderInterface
	tracer          trace.Tracer
}

func NewGetPlaceForecastUseCase(placeProvider place.PlaceProviderInterface, weatherProvider weather.WeatherProviderInterface, tracer trace.Tracer) GetPlaceForecast {
	return GetPlaceForecast{
		placeProvider:   placeProvider,
		weatherProvider: weatherProvider,
		tracer:          tracer,
	}
}

func (uc *GetPlaceForecast) Execute(cep string) (PlaceForecastOutputDTO, error) {
	ctx, span := uc.tracer.Start(context.Background(), "Normalizing Postal Code")
	span.SetAttributes(attribute.String("cep", cep))
	defer span.End()

	normalized, err := entity.NormalizePostalCode(cep)
	outputDTO := PlaceForecastOutputDTO{}

	if err != nil {
		return outputDTO, ErrInvalidInput
	}

	ctx, span = uc.tracer.Start(context.Background(), "Requesting cep place details")
	span.SetAttributes(attribute.String("cep", normalized))
	defer span.End()

	placeDetails, err := uc.placeProvider.GetByCep(normalized, ctx)

	if err != nil || placeDetails.IsValid() != nil {
		return outputDTO, ErrPostalCodeNotFound
	}

	ctx, span = uc.tracer.Start(context.Background(), "Requesting city forecast")
	span.SetAttributes(attribute.String("city", placeDetails.City))
	defer span.End()

	forecast, err := uc.weatherProvider.GetWeather(placeDetails.City, ctx)

	if err != nil {
		return outputDTO, ErrWeatherNotFound
	}

	outputDTO.City = placeDetails.City
	outputDTO.TempC = forecast.TempC
	outputDTO.TempF = forecast.CalculateFahrenheit()
	outputDTO.TempK = forecast.CalculateKelvin()

	return outputDTO, nil
}
