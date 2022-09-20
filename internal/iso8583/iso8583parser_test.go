package iso8583_test

import (
	"api-iso8583-to-JSON/internal/iso8583"
	"testing"

	iso8583_mocks "api-iso8583-to-JSON/internal/iso8583/mocks"

	"github.com/stretchr/testify/assert"
)

var (
	parser = iso8583.NewIso8583Parser()
)

func TestIso8583Parser(t *testing.T) {
	t.Run("Test parse function with 64 position message", func(t *testing.T) {
		parsedMessage, err := parser.Parse(iso8583_mocks.Iso8583Message1Bytes)
		assert.Nil(t, err)
		assert.Equal(t, parsedMessage.Fields[2], iso8583_mocks.Iso8583Message1.Fields[2])
		assert.Equal(t, parsedMessage.Fields[3], iso8583_mocks.Iso8583Message1.Fields[3])
		assert.Equal(t, parsedMessage.Fields[4], iso8583_mocks.Iso8583Message1.Fields[4])
		assert.Equal(t, parsedMessage.Fields[7], iso8583_mocks.Iso8583Message1.Fields[7])
		assert.Equal(t, parsedMessage.Fields[11], iso8583_mocks.Iso8583Message1.Fields[11])
		assert.Equal(t, parsedMessage.Fields[14], iso8583_mocks.Iso8583Message1.Fields[14])
		assert.Equal(t, parsedMessage.Fields[18], iso8583_mocks.Iso8583Message1.Fields[18])
		assert.Equal(t, parsedMessage.Fields[22], iso8583_mocks.Iso8583Message1.Fields[22])
		assert.Equal(t, parsedMessage.Fields[25], iso8583_mocks.Iso8583Message1.Fields[25])
		assert.Equal(t, parsedMessage.Fields[35], iso8583_mocks.Iso8583Message1.Fields[35])
		assert.Equal(t, parsedMessage.Fields[37], iso8583_mocks.Iso8583Message1.Fields[37])
		assert.Equal(t, parsedMessage.Fields[41], iso8583_mocks.Iso8583Message1.Fields[41])
		assert.Equal(t, parsedMessage.Fields[42], iso8583_mocks.Iso8583Message1.Fields[42])
		assert.Equal(t, parsedMessage.Fields[49], iso8583_mocks.Iso8583Message1.Fields[49])
	})

	t.Run("Test parse function with 128 position message", func(t *testing.T) {
		parsedMessage, err := parser.Parse(iso8583_mocks.Iso8583Message2Bytes)
		assert.Nil(t, err)
		assert.Equal(t, parsedMessage.Fields[2], iso8583_mocks.Iso8583Message2.Fields[2])
		assert.Equal(t, parsedMessage.Fields[3], iso8583_mocks.Iso8583Message2.Fields[3])
		assert.Equal(t, parsedMessage.Fields[4], iso8583_mocks.Iso8583Message2.Fields[4])
		assert.Equal(t, parsedMessage.Fields[7], iso8583_mocks.Iso8583Message2.Fields[7])
		assert.Equal(t, parsedMessage.Fields[11], iso8583_mocks.Iso8583Message2.Fields[11])
		assert.Equal(t, parsedMessage.Fields[14], iso8583_mocks.Iso8583Message2.Fields[14])
		assert.Equal(t, parsedMessage.Fields[18], iso8583_mocks.Iso8583Message2.Fields[18])
		assert.Equal(t, parsedMessage.Fields[22], iso8583_mocks.Iso8583Message2.Fields[22])
		assert.Equal(t, parsedMessage.Fields[25], iso8583_mocks.Iso8583Message2.Fields[25])
		assert.Equal(t, parsedMessage.Fields[35], iso8583_mocks.Iso8583Message2.Fields[35])
		assert.Equal(t, parsedMessage.Fields[37], iso8583_mocks.Iso8583Message2.Fields[37])
		assert.Equal(t, parsedMessage.Fields[41], iso8583_mocks.Iso8583Message2.Fields[41])
		assert.Equal(t, parsedMessage.Fields[42], iso8583_mocks.Iso8583Message2.Fields[42])
		assert.Equal(t, parsedMessage.Fields[49], iso8583_mocks.Iso8583Message2.Fields[49])
		assert.Equal(t, parsedMessage.Fields[71], iso8583_mocks.Iso8583Message2.Fields[71])
		assert.Equal(t, parsedMessage.Fields[72], iso8583_mocks.Iso8583Message2.Fields[72])
		assert.Equal(t, parsedMessage.Fields[73], iso8583_mocks.Iso8583Message2.Fields[73])

	})

}
func TestIso8583ParserWithErrors(t *testing.T) {
	t.Run("Test parse function with no message", func(t *testing.T) {
		_, err := parser.Parse([]byte{})
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Invalid Iso8583 Input")
	})

	t.Run("Test parse function with invalid message", func(t *testing.T) {
		_, err := parser.Parse([]byte("123"))
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Error reading MTI")
	})

	t.Run("Test parse function with mti and no data", func(t *testing.T) {
		_, err := parser.Parse([]byte("0200"))
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Error reading first bitmap")
	})

	t.Run("Test parse function with mti and bitmap with no data", func(t *testing.T) {
		_, err := parser.Parse([]byte("02007224448028C08000"))
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Error parsing field 2")
	})

	t.Run("Test parse function with mti and bitmap with no second bitmap", func(t *testing.T) {
		_, err := parser.Parse([]byte("0200F224448028C08000")) // second bitmap present
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Error reading second bitmap")
	})

	t.Run("Test parse function with mti and invalid hex bitmap", func(t *testing.T) {
		_, err := parser.Parse([]byte("0200722444802%8C08000"))
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Error reading first bitmap")
	})
}
