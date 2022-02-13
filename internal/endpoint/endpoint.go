package endpoint

import (
	"api-iso8583-to-JSON/internal/entity"
	"api-iso8583-to-JSON/internal/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeIso8583toJSONEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.Iso8583)

		return svc.Call(ctx, req)
	}
}
