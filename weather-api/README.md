# Weather temperatures API

This is a system that, given a valid 8-digit ZIP code,
identifies the corresponding city and returns the current it's temperature in Celsius, Fahrenheit, and Kelvin.

### Technologies Used

The system is built using the following technologies:

- [Viper](https://github.com/spf13/viper): A Go library for managing application configurations.
- [Wire](https://github.com/google/wire): A dependency injection code generator for Go.
- Native HTTP library for handling requests.

### Customizables

The system is currently using Viacep API to retrieve the city for a zipcode
and WeatherAPI to retrieve the city forescast
Both apis can be replaced by an equivalent through.

The ViaCep can be replaced by any provider as long as that provider implements
the `PlaceProviderInterface` interface found in the `place` package:

```go
type PlaceProviderInterface interface {
	GetByCep(cep string) (entity.Place, error)
}
```

Then, change the wire provider to provide the new place

```go
var setPlaceProviderInterface = wire.NewSet(/* needed arguments*/)
```

The WeatherAPI can be replaced by any provider as long as that provider implements
the `WeatherProviderInterface` interface found in the `weather` package:

```go
type WeatherProviderInterface interface {
	GetWeather(city string) (entity.Weather, error)
}
```

Then, change the wire provider to provide the new weather provider

```go
var setWeatherProviderInterface = wire.NewSet(/* needed arguments*/)
```

# Tests

This command will start the application using Docker Compose and then run the main.go file.

To execute all tests, run the following command:

```bash
make run-tests
```

## Building the project's image

## **Important: Set environment variables in .env before running the project.**

### Production

A docker image ready for production can be built by running

```bash
make build-prod IMAGE_NAME=your_image_name # if empty. weather-api-image:latest will be the default value
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

### Cloud run

The system is deployed on Google Cloud Run and can be accessed at `https://goexpert-cloudrun-weather-api-pwvfjx4fpq-rj.a.run.app`

#### 200

curl -X GET https://goexpert-cloudrun-weather-api-pwvfjx4fpq-rj.a.run.app?cep=89216310

#### 404

curl -X GET https://goexpert-cloudrun-weather-api-pwvfjx4fpq-rj.a.run.app?cep=89216369

#### 422

curl -X GET https://goexpert-cloudrun-weather-api-pwvfjx4fpq-rj.a.run.app?cep=892169

#### Responses

- In case of success:

  - HTTP Code: 200
  - Response Body: `{ "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }`

- In case of failure, if the ZIP code is not valid (with correct format):

  - HTTP Code: 422
  - Message: `invalid zipcode`

- In case of failure, if the ZIP code is not found:
  - HTTP Code: 404
  - Message: `zipcode can not found`
