.PHONY: start run run-tests cep-input weather-api
# Variável para armazenar o parâmetro
CEP_IMAGE_NAME ?= observability-cep-input-image:latest
WEATHER_IMAGE_NAME ?= observability-weather-api-image:latest

start:
	@echo "Starting project..."
	docker-compose up

run: start

build-prod: 
	@echo "Building services..."
	$(MAKE) -C cep-input build-prod IMAGE_NAME=$(CEP_IMAGE_NAME)
	$(MAKE) -C weather-api build-prod IMAGE_NAME=$(WEATHER_IMAGE_NAME)