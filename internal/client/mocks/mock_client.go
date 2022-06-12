package client

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeMockClient() endpoint.Endpoint {
	return func(_ context.Context, req interface{}) (interface{}, error) {
		reqBytes := req.([]byte)
		reqBytes[2] = []byte("1")[0]
		return reqBytes, nil
	}
}
