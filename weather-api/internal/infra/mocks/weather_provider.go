package mocks

import (
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/entity"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/weather"
	"github.com/stretchr/testify/mock"
)

var _ weather.WeatherProviderInterface = (*WeatherProviderMock)(nil)

type WeatherProviderMock struct {
	mock.Mock
}

func NewWeatherProviderMock() *WeatherProviderMock {
	return &WeatherProviderMock{}
}

func (lm *WeatherProviderMock) GetWeather(city string) (entity.Weather, error) {
	args := lm.Called(city)
	return args.Get(0).(entity.Weather), args.Error(1)
}
