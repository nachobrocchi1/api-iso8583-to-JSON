package iso8583

import (
	"api-iso8583-to-JSON/internal/entity"
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
)

const (
	// bitmap lenght
	BitmapBits = 64
)

// BitmapGenerator generate a bitmap from iso8583 struct
type BitmapGenerator interface {
	Generate(iso entity.Iso8583) ([]int, []int, error)
}

type bitmapGenerator struct {
	fieldConfig map[int]iso8583config.FieldConfiguration
}

// NewBitmapGenerator constructor
func NewBitmapGenerator() BitmapGenerator {
	return &bitmapGenerator{
		fieldConfig: iso8583config.GetIsoFieldsConfig(),
	}
}

func (b *bitmapGenerator) Generate(iso entity.Iso8583) ([]int, []int, error) {
	bitmap := make([]int, 2*BitmapBits)
	var bitmap2 []int

	for fieldIndex, _ := range iso.Fields {

		fieldConf := b.fieldConfig[fieldIndex]

		if len(fieldConf.Name) < 1 {
			return nil, nil, iso8583Error("Invalid ISO field")
		}

		if bitmap[0] == 0 && fieldIndex > BitmapBits {
			bitmap[0] = 1
		}

		bitmap[fieldIndex-1] = 1
	}

	if bitmap[0] == 1 {
		bitmap2 = bitmap[BitmapBits:]
	}

	return bitmap[:BitmapBits], bitmap2, nil
}
