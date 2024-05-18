package entity_test

import (
	"testing"

	"github.com/rhbarauna/goexpert-desafio-cloud-run/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCalculateFahrenheit(t *testing.T) {

	weather := entity.Weather{TempC: 20}
	fahrenheit := weather.CalculateFahrenheit()
	assert.Equal(t, 68.0, fahrenheit)

	weather = entity.Weather{TempC: -10}
	fahrenheit = weather.CalculateFahrenheit()
	assert.Equal(t, 14.0, fahrenheit)

	weather = entity.Weather{TempC: 0}
	fahrenheit = weather.CalculateFahrenheit()
	assert.Equal(t, 32.0, fahrenheit)
}

func TestCalculateKelvin(t *testing.T) {
	weather := entity.Weather{TempC: 20}
	kelvin := weather.CalculateKelvin()
	assert.Equal(t, 293.0, kelvin)

	weather = entity.Weather{TempC: -10}
	kelvin = weather.CalculateKelvin()
	assert.Equal(t, 263.0, kelvin)

	weather = entity.Weather{TempC: 0}
	kelvin = weather.CalculateKelvin()
	assert.Equal(t, 273.0, kelvin)
}
