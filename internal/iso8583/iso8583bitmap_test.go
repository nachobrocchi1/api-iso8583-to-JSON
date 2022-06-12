package iso8583_test

import (
	"api-iso8583-to-JSON/internal/iso8583"
	iso8583_mocks "api-iso8583-to-JSON/internal/iso8583/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	bitmapGenerator = iso8583.NewBitmapGenerator()
	isoEncoder      = iso8583.NewIso8583BitmapEncodeDecode()
)

func TestEncode(t *testing.T) {

	t.Run("Valid bitmap encode", func(t *testing.T) {
		bitmap1, _, errBMG := bitmapGenerator.Generate(iso8583_mocks.Iso8583Message1)
		assert.Nil(t, errBMG)
		bitmapBytes, err := isoEncoder.Encode(bitmap1)
		assert.Nil(t, err)
		assert.Equal(t, bitmapM1B1Hex, string(bitmapBytes))
	})

	t.Run("Valid bitmaps encode", func(t *testing.T) {
		bitmap1, bitmap2, errBMG := bitmapGenerator.Generate(iso8583_mocks.Iso8583Message2)
		assert.Nil(t, errBMG)
		bitmapBytes1, err := isoEncoder.Encode(bitmap1)
		assert.Nil(t, err)
		assert.Equal(t, bitmapM2B1Hex, string(bitmapBytes1))
		bitmapBytes2, err := isoEncoder.Encode(bitmap2)
		assert.Nil(t, err)
		assert.Equal(t, bitmapM2B2Hex, string(bitmapBytes2))
	})

	t.Run("Invalid bitmap lenght", func(t *testing.T) {
		_, err := isoEncoder.Encode([]int{2, 4, 5, 7, 11})
		assert.Error(t, err)
		assert.Equal(t, "ISO 8583 Error : Invalid bitmap lenght", err.Error())
	})
}

func TestDecode(t *testing.T) {
	t.Run("Valid bitmap decode", func(t *testing.T) {
		binaryBitmap, err := isoEncoder.Decode([]byte(bitmapM2B1Hex))
		assert.Nil(t, err)
		assert.Equal(t, iso8583_mocks.Iso8583Message2Bitmap1, binaryBitmap)
		binaryBitmap2, err := isoEncoder.Decode([]byte(bitmapM2B2Hex))
		assert.Nil(t, err)
		assert.Equal(t, iso8583_mocks.Iso8583Message2Bitmap2, binaryBitmap2)
	})

	t.Run("Invalid bitmap decode", func(t *testing.T) {
		_, err := isoEncoder.Decode([]byte("GJKASDJKDASODJ"))
		assert.Error(t, err)
		assert.Equal(t, "ISO 8583 Error : Invalid hex", err.Error())
	})
}

var (
	bitmapM1B1Hex = "7224448028C08000"
	bitmapM2B1Hex = "F224448028C08000"
	bitmapM2B2Hex = "0380000000000000"
)
