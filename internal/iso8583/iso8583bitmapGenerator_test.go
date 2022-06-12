package iso8583_test

import (
	"api-iso8583-to-JSON/internal/entity"
	"api-iso8583-to-JSON/internal/iso8583"
	iso8583_mocks "api-iso8583-to-JSON/internal/iso8583/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitmapGenerator(t *testing.T) {
	t.Run("Generates a valid bitmap", func(t *testing.T) {
		bitmapGenerator := iso8583.NewBitmapGenerator()
		bitmap1, bitmap2, err := bitmapGenerator.Generate(iso8583_mocks.Iso8583Message1)
		assert.Nil(t, err)
		assert.Len(t, bitmap1, 64)
		assert.ElementsMatch(t, iso8583_mocks.Iso8583Message1Bitmap, bitmap1)
		assert.Len(t, bitmap2, 0)
	})

	t.Run("Generates valids bitmaps", func(t *testing.T) {
		bitmapGenerator := iso8583.NewBitmapGenerator()
		bitmap1, bitmap2, err := bitmapGenerator.Generate(iso8583_mocks.Iso8583Message2)
		assert.Nil(t, err)
		assert.Len(t, bitmap1, 64)
		assert.ElementsMatch(t, iso8583_mocks.Iso8583Message2Bitmap1, bitmap1)
		assert.Len(t, bitmap2, 64)
		assert.ElementsMatch(t, iso8583_mocks.Iso8583Message2Bitmap2, bitmap2)
	})

	t.Run("Generate fails - invalid ISO field", func(t *testing.T) {
		bitmapGenerator := iso8583.NewBitmapGenerator()
		var isoMessage entity.Iso8583
		isoMessage.Fields = map[int]string{
			999: "wrong field",
		}
		_, _, err := bitmapGenerator.Generate(isoMessage)
		assert.Error(t, err)
		assert.Equal(t, "ISO 8583 Error : Invalid ISO field", err.Error())
	})
}
