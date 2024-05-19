//go:build wireinject
// +build wireinject

package main

import (
	"pos-graduacao/desafios/observabilidade/input/internal/infra/forecast"
	"pos-graduacao/desafios/observabilidade/input/internal/infra/forecast/weather"
	"pos-graduacao/desafios/observabilidade/input/internal/infra/tracing"
	"pos-graduacao/desafios/observabilidade/input/internal/usecase"
	"pos-graduacao/desafios/observabilidade/input/internal/web/handler"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/google/wire"
)

func NewAppTracer() trace.Tracer {
	return otel.Tracer("cep-input")
}

func NewTracing(url, serviceName string) func() {
	return tracing.InitializeTracer(url, serviceName)
}

var setTraceProvider = wire.NewSet(NewAppTracer)

var setForecastProviderInterface = wire.NewSet(
	weather.NewWeatherApi,
	wire.Bind(new(forecast.ForecastProviderInterface), new(*weather.WeatherApi)),
)

func provideGetForecastUC() usecase.GetForecast {
	wire.Build(
		setTraceProvider,
		setForecastProviderInterface,
		usecase.NewGetForecastUseCase,
	)
	return usecase.GetForecast{}
}

func NewCepForecastHandler() handler.GetCepForecastHandler {
	wire.Build(
		provideGetForecastUC,
		handler.NewCepForecastHandler,
	)
	return handler.GetCepForecastHandler{}
}
