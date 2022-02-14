package iso8583

import (
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
)

type Iso8583PositionReader interface {
	ReadPosition(input []byte, start, len int) ([]byte, error)
	ReadMti(input []byte) ([]byte, error)
}

type iso8583PositionIsoReader struct {
	validator Iso8583Validator
}

func NewPositionIsoReader() Iso8583PositionReader {
	return &iso8583PositionIsoReader{
		validator: NewValidator(),
	}
}

func (r *iso8583PositionIsoReader) ReadPosition(input []byte, startPos, fieldLenght int) ([]byte, error) {

	if errField := r.validator.ValidatePosition(input, startPos, fieldLenght); errField != nil {
		return nil, errField
	}

	return input[startPos : startPos+fieldLenght], nil
}

func (m *iso8583PositionIsoReader) ReadMti(input []byte) ([]byte, error) {
	reader := NewPositionIsoReader()
	return reader.ReadPosition(input, iso8583config.MTI_START_POS, iso8583config.MTI_LENGHT)
}
