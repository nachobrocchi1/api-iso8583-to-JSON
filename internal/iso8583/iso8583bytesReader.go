package iso8583

import (
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
	"log"
	"strconv"
)

// Iso8583BytesReader interface for iso reading
type Iso8583BytesReader interface {
	Read(isobytes []byte, startPos, bitmapIndex int) ([]byte, int, error)
	ReadLXVAR(isobytes []byte, startPos, sizeBytes int) ([]byte, int, error)
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
	var fieldConfig iso8583config.FieldConfiguration = r.fieldsConfig[bitmapIndex]

	switch lenType := fieldConfig.LengthType; lenType {
	case iso8583config.LVAR, iso8583config.LLVAR, iso8583config.LLLVAR:
		return r.ReadLXVAR(isobytes, startPos, int(lenType))
	case iso8583config.FIXED:
		return r.readFieldValue(isobytes, startPos, fieldConfig.Length)
	default:
		return nil, startPos, iso8583Error("Invalid field lenght type")
	}

}

func (r *iso8583BytesReader) ReadLXVAR(isobytes []byte, startPos, lenBytes int) ([]byte, int, error) {
	size, err := readLXVARLength(r, isobytes, startPos, lenBytes)

	if err != nil {
		return nil, startPos, err
	}
	// Se agrega a la posicion starPos sumarle el lenBytes, para que continue leyendo el proximo dato
	if size == 0 {
		return []byte(""), startPos + lenBytes, err
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

	if err == nil {
		size, err = strconv.Atoi(string(sizeAsBytes))
	}

	if err != nil {
		log.Printf("Error leyendo LXVAR: %v", err)
		return startPos, iso8583Error("Invalid field size")
	}

	return size, nil
}
