package iso8583

import (
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
	"fmt"
)

const ()

// Iso8583FieldToString interface to converts an iso field into string
type Iso8583FieldToString interface {
	ToString(fieldIndex int, fieldValue string) (string, error)
}

type iso8583FieldToString struct {
	fieldsConfig   map[int]iso8583config.FieldConfiguration
	fieldValidator Iso8583FieldValidator
}

// NewIso8583FieldStringifier crea un nuevo Iso Field Stringifier
func NewIso8583FieldStringifier() Iso8583FieldToString {
	return &iso8583FieldToString{
		fieldsConfig:   iso8583config.GetIsoFieldsConfig(),
		fieldValidator: NewIso8583FieldValidator(),
	}
}

func (s *iso8583FieldToString) ToString(fieldIndex int, fieldValue string) (string, error) {
	conf := s.fieldsConfig[fieldIndex]
	stringifiedValue := preProcessFixedValue(conf, fieldValue)

	if err := s.fieldValidator.Validate(fieldIndex, stringifiedValue); err != nil {
		return "", err
	}

	stringifiedValue = processNonFixedValue(conf, stringifiedValue)

	return stringifiedValue, nil
}

func preProcessFixedValue(conf iso8583config.FieldConfiguration, fieldValue string) string {
	if iso8583config.FIXED != conf.LengthType {
		return fieldValue
	}

	if iso8583config.N == conf.FieldType {
		return fmt.Sprintf("%0*s", conf.Length, fieldValue)
	}

	return fmt.Sprintf("%*s", conf.Length, fieldValue)
}

func processNonFixedValue(conf iso8583config.FieldConfiguration, fieldValue string) string {

	if iso8583config.FIXED == conf.LengthType {
		return fieldValue
	}

	lenbytes := getLenbytesByFieldLengthType(conf)

	fieldLen := len(fieldValue)

	fieldPrefix := fmt.Sprintf("%0*d", lenbytes, fieldLen)

	return fmt.Sprintf("%v%v", fieldPrefix, fieldValue)
}

func getLenbytesByFieldLengthType(conf iso8583config.FieldConfiguration) int {
	switch conf.LengthType {
	case iso8583config.LLVAR:
		return iso8583config.LLVARBytes
	case iso8583config.LLLVAR:
		return iso8583config.LLLVARBytes
	default:
		return iso8583config.LVARBytes
	}
}
