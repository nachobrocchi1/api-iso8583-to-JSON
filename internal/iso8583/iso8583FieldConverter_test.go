package iso8583_test

import (
	"api-iso8583-to-JSON/internal/iso8583"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	fieldConverter = iso8583.NewIso8583FieldConverter()
)

func TestIsoFieldConverter(t *testing.T) {
	t.Run("Convert Field 2", func(t *testing.T) {
		fieldValue, err := fieldConverter.ToISOField(2, "0123456789")
		assert.Nil(t, err)
		assert.Equal(t, "100123456789", fieldValue)
	})

	t.Run("Convert Field 2 with error", func(t *testing.T) {
		_, err := fieldConverter.ToISOField(2, "01234567890123456789")
		assert.Error(t, err)
	})

	t.Run("Convert Field 3", func(t *testing.T) {
		fieldValue, err := fieldConverter.ToISOField(3, "123456")
		assert.Nil(t, err)
		assert.Equal(t, "123456", fieldValue)
	})

	t.Run("Convert Field 3 with zeros", func(t *testing.T) {
		fieldValue, err := fieldConverter.ToISOField(3, "12345")
		assert.Nil(t, err)
		assert.Equal(t, "012345", fieldValue)
	})

	t.Run("Convert Field 49", func(t *testing.T) {
		fieldValue, err := fieldConverter.ToISOField(49, "840")
		assert.Nil(t, err)
		assert.Equal(t, "840", fieldValue)
	})
}
