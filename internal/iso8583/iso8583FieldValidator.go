package iso8583

import (
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
	"fmt"
	"log"
	"regexp"
)

const (
	InvalidFieldLength = iso8583FieldValidatorError("Invalid Iso8583 Field Lenght")
	InvalidFieldFormat = iso8583FieldValidatorError("Invalid Iso8583 Field Format")
)

type iso8583FieldValidatorError string

func (e iso8583FieldValidatorError) Error() string {
	return fmt.Sprintf("Field validation Error [%s]", string(e))
}

var (
	alphanumeric                      = regexp.MustCompile("^[\\w\\s\\*\\-]+$")
	numeric                           = regexp.MustCompile("^\\s*\\d+$|^\\s*$")
	alphanumericWithSpecialCharacters = regexp.MustCompile(`[\w\W]+`)
)

// Iso8583FieldValidator interfaz para validacion de campos ISO8383
type Iso8583FieldValidator interface {
	Validate(fieldIndex int, value string) error
}

type iso8583FieldValidator struct {
	fieldsConfig map[int]iso8583config.FieldConfiguration
}

// NewIso8583FieldValidator constructor para validador de campos ISO8583
func NewIso8583FieldValidator() Iso8583FieldValidator {
	return &iso8583FieldValidator{fieldsConfig: iso8583config.GetIsoFieldsConfig()}
}

func (v *iso8583FieldValidator) Validate(fieldIndex int, value string) (err error) {

	defer func() {
		if err != nil {
			log.Printf("Error reading field: %d value: %q", fieldIndex, value)
		}
	}()

	conf := v.fieldsConfig[fieldIndex]

	if len(conf.Name) < 1 {
		err = InvalidIsoFieldIndex
		return
	}

	if err = validateLength(conf, value); err != nil {
		return
	}

	if len(value) == 0 {
		return
	}

	if err = validateRegex(conf, []byte(value)); err != nil {
		return
	}

	return
}

func validateLength(conf iso8583config.FieldConfiguration, value string) error {
	switch lenType := conf.LengthType; lenType {
	case iso8583config.FIXED:
		if len(value) != conf.Length {
			return InvalidFieldLength
		}
	case iso8583config.LVAR, iso8583config.LLVAR, iso8583config.LLLVAR:
		if len(value) > conf.Length {
			return InvalidFieldLength
		}
	}

	return nil
}

func validateRegex(conf iso8583config.FieldConfiguration, value []byte) error {
	var regex *regexp.Regexp

	switch conf.FieldType {
	case iso8583config.N:
		regex = numeric
	case iso8583config.AN:
		regex = alphanumeric
	case iso8583config.ANS:
		regex = alphanumericWithSpecialCharacters
	case iso8583config.Z:
		return nil // investigar ISO 4909 y en ISO 7813
	}

	if !regex.Match(value) {
		return InvalidFieldFormat
	}

	return nil
}
