package forecast

import (
	"context"
	"pos-graduacao/desafios/observabilidade/input/internal/entity"
)

type ForecastProviderInterface interface {
	GetForecast(cep string, ctx context.Context) (entity.Forecast, error)
}
