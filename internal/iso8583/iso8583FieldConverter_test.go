package iso8583_test

import (
	"api-iso8583-to-JSON/internal/iso8583"
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	fieldConverter = iso8583.NewIso8583FieldConverter(iso8583config.GetIsoFieldsConfig())
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

	t.Run("Convert Field 34", func(t *testing.T) {
		fieldValue, err := fieldConverter.ToISOField(34, "01234567890123456789")
		assert.Nil(t, err)
		assert.Equal(t, "02001234567890123456789", fieldValue)
	})

	t.Run("Convert Field LVAR", func(t *testing.T) {
		fakeFieldConverter := iso8583.NewIso8583FieldConverter(map[int]iso8583config.FieldConfiguration{
			999: {
				Name:       "fake field",
				FieldType:  iso8583config.AN,
				LengthType: iso8583config.LVAR,
				Length:     10,
			},
		})
		fieldValue, err := fakeFieldConverter.ToISOField(999, "012345678")
		assert.Nil(t, err)
		assert.Equal(t, "9012345678", fieldValue)
	})
}
