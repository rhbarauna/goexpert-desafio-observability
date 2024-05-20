//go:build wireinject
// +build wireinject

package main

import (
	"path/filepath"
	"pos-graduacao/desafios/observabilidade/input/configs"
	"pos-graduacao/desafios/observabilidade/input/internal/infra/forecast"
	"pos-graduacao/desafios/observabilidade/input/internal/infra/forecast/weather"
	"pos-graduacao/desafios/observabilidade/input/internal/infra/tracing"
	"pos-graduacao/desafios/observabilidade/input/internal/usecase"
	"pos-graduacao/desafios/observabilidade/input/internal/web/handler"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/google/wire"
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

func NewAppTracer(config *configs.Config) trace.Tracer {
	return otel.Tracer(config.SERVICE_NAME)
}

func NewTracing() func() {
	wire.Build(
		provideConfig,
		tracing.InitializeTracer,
	)
	return func() {}
}

var setTraceProvider = wire.NewSet(NewAppTracer)

var setForecastProviderInterface = wire.NewSet(
	weather.NewWeatherApi,
	wire.Bind(new(forecast.ForecastProviderInterface), new(*weather.WeatherApi)),
)

func provideGetForecastUC() usecase.GetForecast {
	wire.Build(
		provideConfig,
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
