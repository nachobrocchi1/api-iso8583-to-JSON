package iso8583_test

import (
	"api-iso8583-to-JSON/internal/entity"
	"api-iso8583-to-JSON/internal/iso8583"
	"testing"

	iso8583_mocks "api-iso8583-to-JSON/internal/iso8583/mocks"

	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"

	"github.com/stretchr/testify/assert"
)

var (
	unparser = iso8583.NewIso8583Unparser(iso8583config.GetIsoFieldsConfig())
)

func TestIso8583Unparser(t *testing.T) {
	t.Run("Test unparse function with 64 position message", func(t *testing.T) {
		unparsedMessage, err := unparser.Unparse(iso8583_mocks.Iso8583Message1)
		assert.Nil(t, err)
		assert.Equal(t, iso8583_mocks.Iso8583Message1Bytes, unparsedMessage)
	})
	t.Run("Test unparse function with 128 position message", func(t *testing.T) {
		unparsedMessage, err := unparser.Unparse(iso8583_mocks.Iso8583Message2)
		assert.Nil(t, err)
		assert.Equal(t, iso8583_mocks.Iso8583Message2Bytes, unparsedMessage)
	})

}
func TestIso8583UnparserWithErrors(t *testing.T) {
	t.Run("Test unparse function field generation error", func(t *testing.T) {
		var isoMessage entity.Iso8583
		isoMessage.Fields = map[int]string{
			3: "wrong field",
		}
		_, err := unparser.Unparse(isoMessage)
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Error unparsing field 3 , Cause: ISO 8583 Field validation Error : Invalid Iso8583 Field Lenght")
	})

	t.Run("Test unparse function bitmap generation error", func(t *testing.T) {
		var isoMessage entity.Iso8583
		isoMessage.Fields = map[int]string{
			999: "wrong field",
		}
		_, err := unparser.Unparse(isoMessage)
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Invalid ISO field 999")
	})

}
