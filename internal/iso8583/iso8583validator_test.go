package iso8583_test

import (
	"api-iso8583-to-JSON/internal/iso8583"
	iso8583_mocks "api-iso8583-to-JSON/internal/iso8583/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	validator = iso8583.NewValidator()
)

func TestIso8583ValidatorTransaction(t *testing.T) {
	t.Run("Validate transaction", func(t *testing.T) {
		err := validator.ValidateTransaction(iso8583_mocks.Iso8583Message1Bytes)
		assert.Nil(t, err)
	})
	t.Run("Validate transaction error", func(t *testing.T) {
		err := validator.ValidateTransaction([]byte{})
		assert.Error(t, err)
	})
}

func TestIso8583ValidatorPosition(t *testing.T) {
	t.Run("Validate position", func(t *testing.T) {
		err := validator.ValidatePosition(iso8583_mocks.Iso8583Message1Bytes, 20, 19)
		assert.Nil(t, err)
	})
	t.Run("Validate position field index error", func(t *testing.T) {
		err := validator.ValidatePosition(iso8583_mocks.Iso8583Message1Bytes, 200, 19)
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Validation Error : Invalid Iso8583 Field Index")
	})

	t.Run("Validate position no input error", func(t *testing.T) {
		err := validator.ValidatePosition([]byte{}, 200, 19)
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Invalid Iso8583 Input")
	})

	t.Run("Validate position out of range error", func(t *testing.T) {
		err := validator.ValidatePosition(iso8583_mocks.Iso8583Message1Bytes, 132, 4)
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Validation Error : Invalid Iso8583 Field Out Of Range")
	})
}
