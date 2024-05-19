package place

import (
	"context"

	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/entity"
)

type PlaceProviderInterface interface {
	GetByCep(cep string, ctx context.Context) (entity.Place, error)
}
