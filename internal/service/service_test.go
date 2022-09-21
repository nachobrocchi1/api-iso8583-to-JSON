package service_test

import (
	mock_client "api-iso8583-to-JSON/internal/client/mocks"
	"api-iso8583-to-JSON/internal/entity"
	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"
	iso8583_mocks "api-iso8583-to-JSON/internal/iso8583/mocks"
	"errors"

	"github.com/stretchr/testify/assert"

	"api-iso8583-to-JSON/internal/service"
	"context"
	"testing"

	"github.com/go-kit/log"
)

var (
	svc = service.NewService(log.NewNopLogger(), mock_client.MakeMockClient(), iso8583config.GetIsoFieldsConfig())
)

func TestIso8583ValidatorTransaction(t *testing.T) {
	t.Run("test service call", func(t *testing.T) {
		response, err := svc.Call(context.TODO(), iso8583_mocks.Iso8583Message1)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, "0210", response.Mti)
	})

	t.Run("test service client error", func(t *testing.T) {
		noclient_svc := service.NewService(log.NewNopLogger(), func(_ context.Context, req interface{}) (interface{}, error) {
			return nil, errors.New("an error")
		}, iso8583config.GetIsoFieldsConfig())
		response, err := noclient_svc.Call(context.TODO(), iso8583_mocks.Iso8583Message1)
		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("test service unparse request error", func(t *testing.T) {
		response, err := svc.Call(context.TODO(), entity.Iso8583{Mti: "123", Fields: map[int]string{3: "wrong field"}})
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Error unparsing field 3 , Cause: ISO 8583 Field validation Error : Invalid Iso8583 Field Lenght")
		assert.Nil(t, response)
	})

	t.Run("test service parse response error", func(t *testing.T) {
		invalidresponse_service := service.NewService(log.NewNopLogger(), mock_client.MakeMockInvalidResponseClient(), iso8583config.GetIsoFieldsConfig())
		response, err := invalidresponse_service.Call(context.TODO(), iso8583_mocks.Iso8583Message1)
		assert.Error(t, err)
		assert.EqualError(t, err, "ISO 8583 Error : Error reading first bitmap")
		assert.Nil(t, response)
	})
}
