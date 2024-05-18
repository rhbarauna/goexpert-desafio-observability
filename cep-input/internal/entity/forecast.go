package entity

import (
	"errors"
	"unicode"
)

type Forecast struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

var (
	ErrInvalidInputType      = errors.New("the input value must be a string")
	ErrInvalidInputMinLength = errors.New("the input value must have more than 8 digits")
)

func NormalizePostalCode(postalCode string) (string, error) {
	var normalized_cep string

	for _, char := range postalCode {
		if unicode.IsDigit(char) {
			normalized_cep += string(char)
		}
	}

	if normalized_cep == "" {
		return "", ErrInvalidInputType
	}

	if len(normalized_cep) != 8 {
		return "", ErrInvalidInputMinLength
	}

	return normalized_cep, nil
}
