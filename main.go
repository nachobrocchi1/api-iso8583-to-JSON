package main

import (
	"api-iso8583-to-JSON/config"
	"api-iso8583-to-JSON/internal/client"
	"api-iso8583-to-JSON/internal/endpoint"
	"api-iso8583-to-JSON/internal/service"
	"api-iso8583-to-JSON/internal/transport"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func main() {
	config := config.GetApiConfig()
	logger := loggerConfiguration(config)

	clientEp := client.MakeClient(config.BackendURI, config.ClientTimeout)
	svc := service.NewService(logger, clientEp)
	ep := endpoint.MakeIso8583toJSONEndpoint(svc)
	r := transport.NewHttpServer(ep, config.Path, logger)
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
		logger.Log("defaulting to port %s", port)
	}
	logger.Log("Api listening at port", port)
	logger.Log("err", http.ListenAndServe(port, r))
}

func loggerConfiguration(config config.ApiConfig) (logger log.Logger) {
	logger = log.NewJSONLogger(os.Stderr)
	//logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	//logger = log.With(logger, "listen", config.Port, "caller", log.DefaultCaller)
	switch config.LogLevel {
	case "debug":
		logger = level.NewFilter(logger, level.AllowDebug())
	case "info":
		logger = level.NewFilter(logger, level.AllowInfo())
	}
	return logger
}
