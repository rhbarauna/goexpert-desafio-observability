# OBSERVABILITY

This is a system that, given a valid 8-digit ZIP code,
identifies the corresponding city and returns it's Name and current temperature in Celsius, Fahrenheit, and Kelvin.

### Technologies Used

The system is built using the following technologies:

- [Viper](https://github.com/spf13/viper): A Go library for managing application configurations.
- [Wire](https://github.com/google/wire): A dependency injection code generator for Go.
- [Zipkin](https://zipkin.io/): A distributed tracing system that helps gather timing data needed to troubleshoot latency problems in service architectures.
- Native HTTP library for handling requests.

### Customizables

The weather-api service can be customized.
Access https://github.com/rhbarauna/goexpert-desafio-cloud-run for more informations.

## Building the project's image

## **Important: Set weather-api environment variables in /weather-api/cmd/.env before running the project.**

## **Important: Set cep-input environment variables in /cep-input/cmd/.env before running the project.**

### Production

A docker image ready for production can be built by running

```bash
make build-prod CEP_IMAGE_NAME=your_cep_image_name WEATHER_IMAGE_NAME=your_weather_image # if empty. observability-cep-input-image:latest and observability-weather-api-image:latest will be the default values
```

### Development

The system can be tested via a http file contained at /api/get_temperatures.http
OR use an HTTP client like curl or Postman or a Rest Client.

### Docker

```bash
docker-compose up
# OR
make run
# OR
make start
```

#### 200

curl -X GET http://localhost:8080?cep=89216310

#### 404

curl -X GET http://localhost:8080?cep=89216369

#### 422

curl -X GET http://localhost:8080?cep=892169

#### Responses

- In case of success:

  - HTTP Code: 200
  - Response Body: `{ "city":"Joinville" ,"temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }`

- In case of failure, if the ZIP code is not valid (with correct format):

  - HTTP Code: 422
  - Message: `invalid zipcode`

- In case of failure, if the ZIP code is not found:
  - HTTP Code: 404
  - Message: `zipcode can not found`
