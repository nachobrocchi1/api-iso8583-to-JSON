package mocks

import (
	"context"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

func MakeMockClient() endpoint.Endpoint {
	return func(_ context.Context, req interface{}) (interface{}, error) {
		var builder strings.Builder
		reqBytes := req.([]byte)[:4]
		switch string(reqBytes) {
		case "0200":
			builder.Write([]byte("0210"))
			builder.Write(req.([]byte)[4:])
			return []byte(builder.String()), nil
		}
		return req, nil
	}
}
