//go:build wireinject
// +build wireinject

package main

import (
	"pos-graduacao/desafios/observabilidade/input/internal/infra/forecast"
	"pos-graduacao/desafios/observabilidade/input/internal/infra/forecast/weather"
	"pos-graduacao/desafios/observabilidade/input/internal/usecase"
	"pos-graduacao/desafios/observabilidade/input/internal/web/handler"

	"github.com/google/wire"
)

var setForecastProviderInterface = wire.NewSet(
	weather.NewWeatherApi,
	wire.Bind(new(forecast.ForecastProviderInterface), new(*weather.WeatherApi)),
)

func provideGetForecastUC() usecase.GetForecast {
	wire.Build(
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
