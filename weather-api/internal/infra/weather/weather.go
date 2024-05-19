package weather

import (
	"context"

	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/entity"
)

type WeatherProviderInterface interface {
	GetWeather(city string, ctx context.Context) (entity.Weather, error)
}
