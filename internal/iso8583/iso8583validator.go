package iso8583

import (
	"fmt"
)

const (
	InvalidIsoFieldIndex = iso8583ValidatorError("Invalid Iso8583 Field Index")
)

type iso8583ValidatorError string

func (e iso8583ValidatorError) Error() string {
	return fmt.Sprintf("Validation Error [%s]", string(e))
}

type Iso8583Validator interface {
	ValidateTransaction(input []byte) error
	ValidatePosition(input []byte, startPos, fieldLenght int) error
}

type byteIso8583Validator struct{}

//Constructor de IsoValidator
func NewValidator() Iso8583Validator {
	return &byteIso8583Validator{}
}

func (v *byteIso8583Validator) ValidateTransaction(input []byte) error {
	inputLen := len(input)
	if inputLen < 1 {
		return InvalidIso8583Input
	}
	return nil
}

func (v *byteIso8583Validator) ValidatePosition(input []byte, startPos, fieldLenght int) error {
	inputLen := len(input)
	if inputLen < 1 {
		return InvalidIso8583Input
	}

	if startPos < 0 || startPos >= inputLen {
		return InvalidIsoFieldIndex
	}

	endpos := startPos + fieldLenght

	if endpos > inputLen || startPos > endpos {
		return InvalidIsoFieldIndex
	}

	return nil
}
