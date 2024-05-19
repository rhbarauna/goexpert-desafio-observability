//go:build wireinject
// +build wireinject

package main

import (
	"path/filepath"
	"runtime"

	"github.com/google/wire"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/configs"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/place"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/place/viacep"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/tracing"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/weather"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/weather/weatherapi"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/usecase"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/web/handler"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func provideConfig() *configs.Config {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("Erro ao obter informações do arquivo.")
	}
	goDir := filepath.Dir(currentFile)

	config, err := configs.LoadConfig(goDir)
	if err != nil {
		panic(err)
	}
	return config
}

func NewAppTracer() trace.Tracer {
	return otel.Tracer("weather-api")
}

func NewTracing(url, serviceName string) func() {
	return tracing.InitializeTracer(url, serviceName)
}

var setTraceProvider = wire.NewSet(NewAppTracer)

var setPlaceProviderInterface = wire.NewSet(
	viacep.NewViaCep,
	wire.Bind(new(place.PlaceProviderInterface), new(*viacep.ViaCep)),
)

var setWeatherProviderInterface = wire.NewSet(
	weatherapi.NewWeatherAPI,
	wire.Bind(new(weather.WeatherProviderInterface), new(*weatherapi.WeatherApi)),
)

func provideGetPlaceForecastUC() usecase.GetPlaceForecast {
	wire.Build(
		provideConfig,
		setTraceProvider,
		setPlaceProviderInterface,
		setWeatherProviderInterface,
		usecase.NewGetPlaceForecastUseCase,
	)
	return usecase.GetPlaceForecast{}
}

func NewGetPlaceTemperaturesHandler() handler.GetPlaceTemperaturesHandler {
	wire.Build(
		provideGetPlaceForecastUC,
		handler.NewGetPlaceTemperaturesHandler,
	)
	return handler.GetPlaceTemperaturesHandler{}
}
