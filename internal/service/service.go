package service

import (
	"api-iso8583-to-JSON/internal/entity"
	"api-iso8583-to-JSON/internal/iso8583"
	"context"
	"sort"

	"github.com/go-kit/log/level"

	"github.com/go-kit/kit/endpoint"

	"github.com/go-kit/log"
)

type Service interface {
	Call(ctx context.Context, req entity.Iso8583) (*entity.Iso8583, error)
}

type service struct {
	logger   log.Logger
	unparser iso8583.Iso8583Unparser
	parser   iso8583.Iso8583Parser
	client   endpoint.Endpoint
}

func NewService(l log.Logger, clientEp endpoint.Endpoint) Service {
	return &service{
		logger:   l,
		unparser: iso8583.NewIso8583Unparser(),
		parser:   iso8583.NewIso8583Parser(),
		client:   clientEp,
	}
}

func (s *service) Call(ctx context.Context, req entity.Iso8583) (*entity.Iso8583, error) {
	bytes, errUnparse := s.unparser.Unparse(req)
	level.Debug(s.logger).Log("Unparsed request:", string(bytes))
	if errUnparse != nil {
		level.Debug(s.logger).Log("Unparsing request ERROR: ", errUnparse)
	}
	response, clientError := s.client(ctx, bytes)
	if clientError != nil {
		level.Debug(s.logger).Log("Client ERROR: ", clientError)
		return nil, clientError
	}

	level.Debug(s.logger).Log("Client Response: ", response)

	res, errParse := s.parser.Parse(response.([]byte))
	level.Debug(s.logger).Log("Parsed response :", res)
	if errParse != nil {
		level.Debug(s.logger).Log("Parsing Response ERROR: ", errParse)
	}

	return orderOutput(res), nil
}

func orderOutput(res *entity.Iso8583) *entity.Iso8583 {
	keys := make([]int, 0, len(res.Fields))
	for k := range res.Fields {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	apiResponse := new(entity.Iso8583)
	apiResponse.Mti = res.Mti
	apiResponse.Fields = make(map[int]string, len(keys))

	for k, v := range res.Fields {
		apiResponse.Fields[k] = v
	}
	return apiResponse
}
