package forecast

import (
	"pos-graduacao/desafios/observabilidade/input/internal/entity"
)

type ForecastProviderInterface interface {
	GetForecast(cep string) (entity.Forecast, error)
}
