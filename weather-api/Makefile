.PHONY: start run run-tests
# Variável para armazenar o parâmetro
IMAGE_NAME ?= weather-api-image:latest

start:
	@echo "Starting project..."
	go run cmd/main.go cmd/wire_gen.go

run: start

build-prod: 
	@echo "Building docker image $(IMAGE_NAME)..."
	docker build -t $(IMAGE_NAME) -f Dockerfile.prod .
run-tests:
	go test ./... -v
