package forecast

import (
	"context"
	"errors"
	"pos-graduacao/desafios/observabilidade/input/internal/entity"
)

var (
	ErrZipCodeNotFound = errors.New("can not find zipcode")
)

type ForecastProviderInterface interface {
	GetForecast(cep string, ctx context.Context) (entity.Forecast, error)
}
