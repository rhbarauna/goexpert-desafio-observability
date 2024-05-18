package entity

import (
	"errors"
	"unicode"
)

type Place struct {
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

func (p *Place) IsValid() error {
	if p.City == "" {
		return errors.New("city not found")
	}
	return nil
}

func NormalizePostalCode(postalCode string) (string, error) {
	var normalized_cep string

	for _, char := range postalCode {
		if unicode.IsDigit(char) {
			normalized_cep += string(char)
		}
	}

	if normalized_cep == "" {
		return "", errors.New("empty postal code")
	}

	if len(normalized_cep) != 8 {
		return "", errors.New("invalid postalcode length")
	}

	return normalized_cep, nil
}
