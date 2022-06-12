package main

import (
	"api-iso8583-to-JSON/config"
	"api-iso8583-to-JSON/internal/client"
	mocksclient "api-iso8583-to-JSON/internal/client/mocks"
	"api-iso8583-to-JSON/internal/endpoint"
	"api-iso8583-to-JSON/internal/service"
	mocksservice "api-iso8583-to-JSON/internal/service/mocks"
	"api-iso8583-to-JSON/internal/transport"
	"fmt"
	"net/http"
	"os"
	"time"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {
	cfg := config.GetApiConfig()
	logger := loggerConfiguration(cfg)

	var clientEp kitendpoint.Endpoint
	var svc service.Service
	{
		if cfg.Backend == "mock" {
			clientEp = mocksclient.MakeMockClient()
			svc = mocksservice.NewMockService(logger, clientEp)
			fmt.Println(config.ColorYellow, "USING MOCK CLIENT")
		} else {
			clientEp = client.MakeClient(cfg.Backend, time.Duration(cfg.ClientTimeout))
			svc = service.NewService(logger, clientEp)

		}
	}

	ep := endpoint.MakeIso8583toJSONEndpoint(svc)
	if cfg.Backend == "mock" {
		ep = endpoint.LoggingMockClientEndpointMiddleware(logger)(ep)
	}

	r := transport.NewHttpServer(ep, cfg.Path, logger)
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port
	}
	level.Info(logger).Log("Api listening at port", port)
	level.Info(logger).Log("err", http.ListenAndServe(":"+port, r))
}

func loggerConfiguration(config config.ApiConfig) (logger log.Logger) {
	logger = log.NewJSONLogger(os.Stderr)
	//logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	switch config.LogLevel {
	case "debug":
		logger = level.NewFilter(logger, level.AllowDebug())
	case "info":
		logger = level.NewFilter(logger, level.AllowInfo())
	}
	return logger
}
