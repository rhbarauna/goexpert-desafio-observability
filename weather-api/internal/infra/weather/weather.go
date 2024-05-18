package weather

import "github.com/rhbarauna/goexpert-desafio-cloud-run/internal/entity"

type WeatherProviderInterface interface {
	GetWeather(city string) (entity.Weather, error)
}
