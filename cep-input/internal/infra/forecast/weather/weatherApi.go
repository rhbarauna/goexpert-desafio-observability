package weather

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"pos-graduacao/desafios/observabilidade/input/internal/entity"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type WeatherApi struct {
	httpClient http.Client
	tracer     trace.Tracer
}

func NewWeatherApi(tracer trace.Tracer) *WeatherApi {
	return &WeatherApi{
		httpClient: http.Client{},
		tracer:     tracer,
	}
}

func (wp *WeatherApi) GetForecast(cep string, ctx context.Context) (entity.Forecast, error) {
	ctx, span := wp.tracer.Start(ctx, "call_weather_api")
	span.SetAttributes(attribute.String("cep", cep))
	defer span.End()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://weather-api:8080?cep=%s", cep), nil)
	forecast := entity.Forecast{}

	if err != nil {
		log.Printf("Falha ao montar a requisição ao serviço de forecast. %s\n", err.Error())
		return forecast, err
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
	resp, err := wp.httpClient.Do(req)

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
