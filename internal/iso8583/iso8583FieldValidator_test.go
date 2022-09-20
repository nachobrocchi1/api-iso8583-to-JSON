package iso8583_test

import (
	"api-iso8583-to-JSON/internal/iso8583"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	fieldValidator = iso8583.NewIso8583FieldValidator()
)

func TestIsoFieldValidator(t *testing.T) {
	t.Run("Validate Field 2", func(t *testing.T) {
		err := fieldValidator.Validate(2, "0123456789")
		assert.Nil(t, err)
	})

	t.Run("Validate Field 500 returns error", func(t *testing.T) {
		err := fieldValidator.Validate(500, "0123456789")
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Validation Error : Invalid Iso8583 Field Index")
	})

	t.Run("Validate Field 3 invalid fixed lenght", func(t *testing.T) {
		err := fieldValidator.Validate(3, "0123456789")
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Field validation Error : Invalid Iso8583 Field Lenght")
	})

	t.Run("Validate Field 2 invalid variable lenght", func(t *testing.T) {
		err := fieldValidator.Validate(3, "012345678911121314151617181920")
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Field validation Error : Invalid Iso8583 Field Lenght")
	})

	t.Run("Validate Field 2 empty", func(t *testing.T) {
		err := fieldValidator.Validate(2, "")
		assert.Nil(t, err)
	})

	t.Run("Validate Field 2 invalid field type", func(t *testing.T) {
		err := fieldValidator.Validate(2, "PAN 12345678")
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Field validation Error : Invalid Iso8583 Field Format")
	})
}
