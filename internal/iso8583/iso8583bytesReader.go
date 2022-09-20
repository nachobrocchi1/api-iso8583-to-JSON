package iso8583

import (
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
	"strconv"
)

// Iso8583BytesReader interface for iso reading
type Iso8583BytesReader interface {
	Read(isobytes []byte, startPos, bitmapIndex int) ([]byte, int, error)
}

type iso8583BytesReader struct {
	fieldsConfig   map[int]iso8583config.FieldConfiguration
	positionReader Iso8583PositionReader
}

// NewIso8583BytesReader Iso Bytes Reader
func NewIso8583BytesReader() Iso8583BytesReader {

	return &iso8583BytesReader{
		fieldsConfig:   iso8583config.GetIsoFieldsConfig(),
		positionReader: NewPositionIsoReader(),
	}
}

func (r *iso8583BytesReader) Read(isobytes []byte, startPos, bitmapIndex int) ([]byte, int, error) {
	fieldConfig, err := r.getIsoFieldConfig(bitmapIndex)
	if err != nil {
		return isobytes, 0, err
	}
	switch lenType := fieldConfig.LengthType; lenType {
	case iso8583config.LVAR, iso8583config.LLVAR, iso8583config.LLLVAR:
		return r.readLXVARValue(isobytes, startPos, int(lenType))
	case iso8583config.FIXED:
		return r.readFieldValue(isobytes, startPos, fieldConfig.Length)
	default:
		return nil, startPos, iso8583Error("Invalid field lenght type")
	}
}

func (r *iso8583BytesReader) getIsoFieldConfig(index int) (iso8583config.FieldConfiguration, error) {
	config, ok := r.fieldsConfig[index]
	if ok {
		return config, nil
	}
	return config, iso8583Error("Invalid field position")
}

func (r *iso8583BytesReader) readLXVARValue(isobytes []byte, startPos, lenBytes int) ([]byte, int, error) {
	size, err := readLXVARLength(r, isobytes, startPos, lenBytes)
	if err != nil {
		return nil, startPos, err
	}

	if size == 0 {
		return nil, startPos, iso8583Error("Field size cannot be zero")
	}

	return r.readFieldValue(isobytes, startPos+lenBytes, size)
}

func (r *iso8583BytesReader) readFieldValue(isobytes []byte, startPos, size int) ([]byte, int, error) {
	valbytes, err := r.positionReader.ReadPosition(isobytes, startPos, size)
	if err != nil {
		return nil, startPos, err
	}

	return valbytes, startPos + size, nil
}

func readLXVARLength(r *iso8583BytesReader, isobytes []byte, startPos, lenBytes int) (int, error) {
	size := 0
	sizeAsBytes, err := r.positionReader.ReadPosition(isobytes, startPos, lenBytes)
	if err != nil {
		return startPos, err
	}

	size, _ = strconv.Atoi(string(sizeAsBytes))

	return size, nil
}
