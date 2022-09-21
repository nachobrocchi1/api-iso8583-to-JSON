package iso8583_test

import (
	"api-iso8583-to-JSON/internal/iso8583"
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	bytes       = []byte("01007224448028C080001643211234432112340000000000000123000304054133001205020553990220010231231233220630500001429110001        1001001840")
	bytesReader = iso8583.NewIso8583BytesReader(iso8583config.GetIsoFieldsConfig())
)

func TestByteReader(t *testing.T) {
	t.Run("Read fields from bytes", func(t *testing.T) {

		//Read field 2 - PAN
		field2, pos3, err2 := bytesReader.Read(bytes, 20, 2)
		assert.Nil(t, err2)
		assert.Equal(t, 38, pos3)
		assert.Equal(t, "4321123443211234", string(field2))

		//Read field 3 - Processing code
		field3, pos4, err3 := bytesReader.Read(bytes, pos3, 3)
		assert.Nil(t, err3)
		assert.Equal(t, 44, pos4)
		assert.Equal(t, "000000", string(field3))
	})

}

func TestByteReaderErrors(t *testing.T) {
	t.Run("Read wrong position field", func(t *testing.T) {
		_, _, err := bytesReader.Read(bytes, 20, 999)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Invalid field position")
	})

	t.Run("Read field in wrong position", func(t *testing.T) {

		//Read field 2 - PAN
		_, _, err := bytesReader.Read(bytes, 38, 2)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Field size cannot be zero")
	})

	t.Run("Read LVAR field in invalid position", func(t *testing.T) {

		//Read field 2 - PAN
		_, _, err := bytesReader.Read(bytes, -1, 2)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "ISO 8583 Validation Error : Invalid Iso8583 Field Index")
	})

	t.Run("Read FIXED field in invalid position", func(t *testing.T) {

		//Read field 3 - Processing code
		_, _, err := bytesReader.Read(bytes, -1, 3)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "ISO 8583 Validation Error : Invalid Iso8583 Field Index")
	})
}

func Test4ByteReaderErrors(t *testing.T) {

}
