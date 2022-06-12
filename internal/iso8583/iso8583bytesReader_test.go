package iso8583_test

import (
	"api-iso8583-to-JSON/internal/iso8583"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	bytes       = []byte("01007224448028C080001643211234432112340000000000000123000304054133001205020553990220010231231233220630500001429110001        1001001840")
	bytesReader = iso8583.NewIso8583BytesReader()
)

func TestByteReader(t *testing.T) {
	t.Run("Read fields from bytes", func(t *testing.T) {

		//Read field 2 - PAN
		field2, pos3, err2 := bytesReader.Read(bytes, 20, 2)
		assert.Nil(t, err2)
		assert.Equal(t, 38, pos3)
		assert.Equal(t, "4321123443211234", string(field2))

		//Read field 3 - Processing code
		field3, pos4, err3 := bytesReader.Read(bytes, 20, 2)
		assert.Nil(t, err3)
		assert.Equal(t, 38, pos4)
		assert.Equal(t, "4321123443211234", string(field3))
	})
}
