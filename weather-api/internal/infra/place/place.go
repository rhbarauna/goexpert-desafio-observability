package place

import "github.com/rhbarauna/goexpert-desafio-cloud-run/internal/entity"

type PlaceProviderInterface interface {
	GetByCep(cep string) (entity.Place, error)
}
