package endpoint

import (
	"api-iso8583-to-JSON/internal/entity"
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func LoggingMockClientEndpointMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {

			var response mockResponse

			res, err := next(ctx, request)

			r := res.(*entity.Iso8583)
			response.Mti = r.Mti
			response.Fields = r.Fields
			response.Request = r.Request
			response.Response = r.Response

			return response, err
		}
	}
}

type mockResponse struct {
	Mti      string         `json:"mti" validate:"required"`
	Fields   map[int]string `json:"fields"`
	Request  string         `json:"requestBytes"`
	Response string         `json:"responseBytes"`
}
