package iso8583

import (
	"api-iso8583-to-JSON/internal/entity"
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
	"fmt"
	"strings"
)

// Iso8583Unparser intefaz de "des" parseo de ISO desde estructura
type Iso8583Unparser interface {
	Unparse(iso entity.Iso8583) ([]byte, error)
}

type iso8583Unparser struct {
	isoFieldConverter Iso8583FieldConverter
	bitmapGenerator   BitmapGenerator
	bitmapEndocer     IsoBitmap
}

// NewIso8583Unparser crea un nuevo unparser Iso8583
func NewIso8583Unparser(config map[int]iso8583config.FieldConfiguration) Iso8583Unparser {
	return &iso8583Unparser{
		isoFieldConverter: NewIso8583FieldConverter(config),
		bitmapGenerator:   NewBitmapGenerator(),
		bitmapEndocer:     NewIso8583BitmapEncodeDecode(),
	}
}

func (u *iso8583Unparser) Unparse(iso entity.Iso8583) ([]byte, error) {
	bitmap1, bitmap2, err := u.bitmapGenerator.Generate(iso)
	if err != nil {
		return nil, err
	}

	bitmapsHex := u.getBitmapsHex(bitmap1, bitmap2)
	stringifiedFields, err := u.isoFieldsToString(iso, bitmap1, bitmap2)
	if err != nil {
		return nil, err
	}

	var unparsedIsoBuilder strings.Builder
	unparsedIsoBuilder.WriteString(iso.Mti)
	unparsedIsoBuilder.WriteString(bitmapsHex)
	unparsedIsoBuilder.WriteString(stringifiedFields)

	unparsedIso := unparsedIsoBuilder.String()

	return []byte(unparsedIso), nil
}

func (u *iso8583Unparser) isoFieldsToString(iso entity.Iso8583, bitmap1 []int, bitmap2 []int) (string, error) {
	var fields strings.Builder

	joinedBitmaps := append(bitmap1, bitmap2...)

	for bitIndex, bit := range joinedBitmaps {
		if bitIndex == 0 || bit == 0 {
			continue
		}

		fieldIndex := bitIndex + 1
		fieldValue := iso.Fields[fieldIndex]
		stringifiedField, err := u.isoFieldConverter.ToISOField(fieldIndex, fieldValue)

		if err != nil {
			return "", iso8583Error(fmt.Sprintf("Error unparsing field %d , Cause: %s", fieldIndex, err.Error()))
		}

		fields.WriteString(stringifiedField)
	}

	return fields.String(), nil
}

func (u *iso8583Unparser) getBitmapsHex(bitmap1 []int, bitmap2 []int) string {
	var bitmapsHex strings.Builder
	bitmap1Hex, _ := u.bitmapEndocer.Encode(bitmap1)

	bitmapsHex.WriteString(string(bitmap1Hex))

	if bitmap2 != nil {
		bitmap2Hex, _ := u.bitmapEndocer.Encode(bitmap2)

		bitmapsHex.WriteString(string(bitmap2Hex))
	}

	return bitmapsHex.String()
}
