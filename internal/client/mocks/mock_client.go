package mocks

import (
	"api-iso8583-to-JSON/vendor/github.com/go-kit/kit/endpoint"
	"context"
	"strings"
	"time"
)

func MakeMockClient(clientUri string, timeout time.Duration) endpoint.Endpoint {
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
