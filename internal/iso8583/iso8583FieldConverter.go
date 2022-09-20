package iso8583

import (
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
	"fmt"
)

const ()

// Iso8583FieldConverter interface to converts an iso field into string
type Iso8583FieldConverter interface {
	ToISOField(fieldIndex int, fieldValue string) (string, error)
}

type iso8583FieldConverter struct {
	fieldsConfig   map[int]iso8583config.FieldConfiguration
	fieldValidator Iso8583FieldValidator
}

// NewIso8583FieldConverter constructor Iso Field Converter
func NewIso8583FieldConverter() Iso8583FieldConverter {
	return &iso8583FieldConverter{
		fieldsConfig:   iso8583config.GetIsoFieldsConfig(),
		fieldValidator: NewIso8583FieldValidator(),
	}
}

func (s *iso8583FieldConverter) ToISOField(fieldIndex int, fieldValue string) (string, error) {
	conf := s.fieldsConfig[fieldIndex]
	var stringifiedValue string
	if iso8583config.FIXED == conf.LengthType {
		stringifiedValue = preProcessFixedValue(conf, fieldValue)
	} else {
		stringifiedValue = processNonFixedValue(conf, fieldValue)
	}

	if err := s.fieldValidator.Validate(fieldIndex, stringifiedValue); err != nil {
		return "", err
	}

	return stringifiedValue, nil
}

func preProcessFixedValue(conf iso8583config.FieldConfiguration, fieldValue string) string {
	if iso8583config.N == conf.FieldType {
		return fmt.Sprintf("%0*s", conf.Length, fieldValue)
	}
	return fmt.Sprintf("%*s", conf.Length, fieldValue)
}

func processNonFixedValue(conf iso8583config.FieldConfiguration, fieldValue string) string {
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
