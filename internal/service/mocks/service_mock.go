package mocks

import (
	"api-iso8583-to-JSON/internal/entity"
	"api-iso8583-to-JSON/internal/iso8583"
	"api-iso8583-to-JSON/internal/service"
	"context"

	"github.com/go-kit/log/level"

	"github.com/go-kit/kit/endpoint"

	iso8583config "api-iso8583-to-JSON/internal/iso8583/config"

	"github.com/go-kit/log"
)

type mockservice struct {
	logger   log.Logger
	unparser iso8583.Iso8583Unparser
	parser   iso8583.Iso8583Parser
	client   endpoint.Endpoint
}

func NewMockService(l log.Logger, clientEp endpoint.Endpoint, config map[int]iso8583config.FieldConfiguration) service.Service {
	return &mockservice{
		logger:   l,
		unparser: iso8583.NewIso8583Unparser(config),
		parser:   iso8583.NewIso8583Parser(config),
		client:   clientEp,
	}
}

func (s *mockservice) Call(ctx context.Context, req entity.Iso8583) (*entity.Iso8583, error) {
	level.Debug(s.logger).Log("BEGIN OF SERVICE CALL", "-------------------------------------------------------------------------")
	level.Debug(s.logger).Log("Request", req)
	response := new(entity.Iso8583)
	requestBytes, errUnparse := s.unparser.Unparse(req)
	if errUnparse != nil {
		level.Error(s.logger).Log("Unparsing request ERROR: ", errUnparse)
		return nil, errUnparse
	}
	response.Request = string(requestBytes)
	level.Debug(s.logger).Log("Unparsed request:", string(requestBytes))

	responseBytes, clientError := s.client(ctx, requestBytes)
	if clientError != nil {
		level.Error(s.logger).Log("Client ERROR: ", clientError)
		return nil, clientError
	}
	response.Response = string(responseBytes.([]byte))
	level.Debug(s.logger).Log("Client Response: ", string(responseBytes.([]byte)))

	res, errParse := s.parser.Parse(responseBytes.([]byte))
	if errParse != nil {
		level.Error(s.logger).Log("Parsing Response ERROR: ", errParse)
		return nil, errParse
	}
	response.Mti = res.Mti
	response.Fields = res.Fields
	level.Debug(s.logger).Log("Parsed response :", res)
	level.Debug(s.logger).Log("END OF SERVICE CALL", "-------------------------------------------------------------------------")
	return response, nil
}
