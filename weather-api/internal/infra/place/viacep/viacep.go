package viacep

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/entity"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/place"
)

var _ place.PlaceProviderInterface = (*ViaCep)(nil)

type ViaCep struct {
	httpClient http.Client
}

type ViaCepResp struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func NewViaCep() *ViaCep {
	return &ViaCep{
		httpClient: http.Client{},
	}
}

func (v *ViaCep) GetByCep(cep string) (entity.Place, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep), nil)
	place := entity.Place{}

	if err != nil {
		log.Printf("Falha ao montar a requisição ViaCep. %s\n", err.Error())
		return place, err
	}

	resp, err := v.httpClient.Do(req)

	if err != nil {
		log.Printf("Falha ao executar a requisição ViaCep. %s\n", err.Error())
		return place, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Erro na resposta da API:", resp.Status)
		return place, errors.New("Erro na resposta da API ViaCep")
	}

	var viacepResp ViaCepResp

	err = json.NewDecoder(resp.Body).Decode(&viacepResp)
	if err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		return place, err
	}

	place.City = viacepResp.Localidade
	place.PostalCode = viacepResp.Cep

	return place, nil
}
