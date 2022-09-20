package service

import (
	"api-iso8583-to-JSON/internal/entity"
	"api-iso8583-to-JSON/internal/iso8583"
	"context"

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
	level.Debug(s.logger).Log("BEGIN OF SERVICE CALL", "-------------------------------------------------------------------------")
	level.Debug(s.logger).Log("Request", req)
	var res *entity.Iso8583
	requestBytes, errUnparse := s.unparser.Unparse(req)
	if errUnparse != nil {
		level.Error(s.logger).Log("Unparsing request ERROR: ", errUnparse)
		return nil, errUnparse
	}
	level.Debug(s.logger).Log("Unparsed request:", string(requestBytes))

	responseBytes, clientError := s.client(ctx, requestBytes)
	if clientError != nil {
		level.Error(s.logger).Log("Client ERROR: ", clientError)
		return nil, clientError
	}
	level.Debug(s.logger).Log("Client Response: ", string(responseBytes.([]byte)))

	res, errParse := s.parser.Parse(responseBytes.([]byte))
	if errParse != nil {
		level.Error(s.logger).Log("Parsing Response ERROR: ", errParse)
		return nil, errParse
	}
	level.Debug(s.logger).Log("Parsed response :", res)
	level.Debug(s.logger).Log("END OF SERVICE CALL", "-------------------------------------------------------------------------")
	return res, nil
}
