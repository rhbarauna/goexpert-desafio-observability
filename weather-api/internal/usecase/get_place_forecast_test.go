package usecase_test

import (
	"testing"

	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/entity"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/infra/mocks"
	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetPlaceForecastSuite struct {
	suite.Suite
	usecase.GetPlaceForecast
	mockPlaceProvider   *mocks.PlaceProviderMock
	mockWeatherProvider *mocks.WeatherProviderMock
}

func (suite *GetPlaceForecastSuite) SetupTest() {
	suite.mockPlaceProvider = new(mocks.PlaceProviderMock)
	suite.mockWeatherProvider = new(mocks.WeatherProviderMock)
	suite.GetPlaceForecast = usecase.NewGetPlaceForecastUseCase(suite.mockPlaceProvider, suite.mockWeatherProvider)
}

func (suite *GetPlaceForecastSuite) TestGetPlaceForecast_Execute_Success() {
	suite.mockPlaceProvider.On("GetByCep", "12345678").Return(entity.Place{City: "New York"}, nil)
	suite.mockWeatherProvider.On("GetWeather", "New York").Return(entity.Weather{TempC: 20}, nil)

	outputDTO, err := suite.GetPlaceForecast.Execute("12345678")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 20.0, outputDTO.TempC)
	assert.Equal(suite.T(), 68.0, outputDTO.TempF)
	assert.Equal(suite.T(), 293.0, outputDTO.TempK)

	suite.mockPlaceProvider.AssertExpectations(suite.T())
	suite.mockWeatherProvider.AssertExpectations(suite.T())
}

func (suite *GetPlaceForecastSuite) TestGetPlaceForecast_Execute_InvalidPostalCode() {
	outputDTO, err := suite.GetPlaceForecast.Execute("invalid-cep")

	assert.EqualError(suite.T(), err, usecase.ErrInvalidInput.Error())
	assert.Empty(suite.T(), outputDTO)

	suite.mockPlaceProvider.AssertNotCalled(suite.T(), "GetByCep")
	suite.mockWeatherProvider.AssertNotCalled(suite.T(), "GetWeather")
}

func (suite *GetPlaceForecastSuite) TestGetPlaceForecast_Execute_PostalCodeNotFound() {
	suite.mockPlaceProvider.On("GetByCep", "12345678").Return(entity.Place{}, usecase.ErrPostalCodeNotFound)

	outputDTO, err := suite.GetPlaceForecast.Execute("12345678")

	assert.EqualError(suite.T(), err, usecase.ErrPostalCodeNotFound.Error())
	assert.Empty(suite.T(), outputDTO)

	suite.mockPlaceProvider.AssertExpectations(suite.T())
	suite.mockWeatherProvider.AssertNotCalled(suite.T(), "GetWeather")
}

func (suite *GetPlaceForecastSuite) TestGetPlaceForecast_Execute_WeatherNotFound() {
	suite.mockPlaceProvider.On("GetByCep", "12345678").Return(entity.Place{City: "New York"}, nil)
	suite.mockWeatherProvider.On("GetWeather", "New York").Return(entity.Weather{}, usecase.ErrWeatherNotFound)

	outputDTO, err := suite.GetPlaceForecast.Execute("12345678")

	assert.EqualError(suite.T(), err, usecase.ErrWeatherNotFound.Error())
	assert.Empty(suite.T(), outputDTO)

	suite.mockPlaceProvider.AssertExpectations(suite.T())
	suite.mockWeatherProvider.AssertExpectations(suite.T())
}

func TestGetPlaceForecastSuite(t *testing.T) {
	suite.Run(t, new(GetPlaceForecastSuite))
}
