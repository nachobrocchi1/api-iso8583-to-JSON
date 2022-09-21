package iso8583

import (
	"api-iso8583-to-JSON/internal/entity"
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Iso8583Parser interface {
	Parse(isoBytes []byte) (*entity.Iso8583, error)
}

type iso8583Parser struct {
	validator         Iso8583Validator
	bitmapDecoder     IsoBitmap
	positionReader    Iso8583PositionReader
	bytesReader       Iso8583BytesReader
	fieldCofiguration map[int]iso8583config.FieldConfiguration
}

func NewIso8583Parser(config map[int]iso8583config.FieldConfiguration) Iso8583Parser {
	return &iso8583Parser{
		validator:         NewValidator(),
		positionReader:    NewPositionIsoReader(),
		bitmapDecoder:     NewIso8583BitmapEncodeDecode(),
		fieldCofiguration: config,
		bytesReader:       NewIso8583BytesReader(config),
	}
}

func (p *iso8583Parser) Parse(isoBytes []byte) (*entity.Iso8583, error) {
	iso8583 := &entity.Iso8583{Fields: make(map[int]string)}

	//clean special chars
	isoBytes = cleanTildeChars(isoBytes)

	//ISO tx validation
	if err := p.validator.ValidateTransaction(isoBytes); err != nil {
		return nil, err
	}

	//MTI
	mti, err := p.positionReader.ReadMti(isoBytes)
	if err != nil {
		return nil, err
	}

	iso8583.Mti = string(mti)

	//Bitmap
	bitmap, currentPosition, err := p.readBitMaps(isoBytes)
	if err != nil {
		return nil, err
	}

	for bitIndex, field := range bitmap {
		if bitIndex == 0 {
			continue
		}

		if field == 1 {
			isoFieldIndex := bitIndex + 1
			fieldVal, nextPosition, err := p.bytesReader.Read(isoBytes, currentPosition, isoFieldIndex)

			if err != nil {
				return nil, iso8583Error(fmt.Sprintf("Error parsing field %d", isoFieldIndex))
			}

			currentPosition = nextPosition
			iso8583.Fields[isoFieldIndex] = strings.ReplaceAll(string(fieldVal), " ", "")
		}
	}

	return iso8583, nil
}

func (p *iso8583Parser) readBitMaps(isoBytes []byte) ([]int, int, error) {

	bitmap1, err := p.readBitMap(isoBytes, iso8583config.FirstBitMapStartPos)
	if err != nil {
		return nil, iso8583config.FirstBitMapStartPos, iso8583Error("Error reading first bitmap")
	}

	if bitmap1[0] == 1 {
		bitmap2, err := p.readBitMap(isoBytes, iso8583config.SecondBitMapStartPos)
		if err != nil {
			return nil, iso8583config.SecondBitMapStartPos, iso8583Error("Error reading second bitmap")
		}

		return append(bitmap1, bitmap2...), iso8583config.IsoDataStartPosition, nil
	}

	return bitmap1, iso8583config.SecondBitMapStartPos, nil

}

func (p *iso8583Parser) readBitMap(isoBytes []byte, startPos int) ([]int, error) {
	hexBytes, err := p.positionReader.ReadPosition(isoBytes, startPos, iso8583config.BitmapLen)
	if err != nil {
		return nil, err
	}

	bitmap, err := p.bitmapDecoder.Decode(hexBytes)

	if err != nil {
		return nil, err
	}

	return bitmap, nil
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func cleanTildeChars(isoBytes []byte) []byte {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)

	transformedIso, _, _ := transform.Bytes(t, isoBytes)

	return transformedIso
}
