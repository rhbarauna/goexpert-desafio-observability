package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"pos-graduacao/desafios/observabilidade/input/internal/entity"
)

type WeatherApi struct {
	httpClient http.Client
}

func NewWeatherApi() *WeatherApi {
	return &WeatherApi{
		httpClient: http.Client{},
	}
}

func (fp *WeatherApi) GetForecast(cep string) (entity.Forecast, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://weather-api:8080?cep=%s", cep), nil)
	forecast := entity.Forecast{}

	if err != nil {
		log.Printf("Falha ao montar a requisição ao serviço de forecast. %s\n", err.Error())
		return forecast, err
	}

	resp, err := fp.httpClient.Do(req)

	if err != nil {
		log.Printf("Falha ao executar a requisição ao serviço de forecast. %s\n", err.Error())
		return forecast, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Erro na resposta da API:", resp.Status)
		return forecast, errors.New("Erro na resposta do serviço de forecast")
	}

	err = json.NewDecoder(resp.Body).Decode(&forecast)
	if err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		return forecast, err
	}

	return forecast, nil
}
