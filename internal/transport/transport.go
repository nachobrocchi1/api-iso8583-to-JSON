package transport

import (
	"api-iso8583-to-JSON/internal/entity"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/go-kit/kit/endpoint"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
)

func NewHttpServer(ep endpoint.Endpoint, path string, logger log.Logger) *mux.Router {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}
	handler := httptransport.NewServer(
		ep,
		makeDecodeRequest(logger),
		makeEncodeResponse(logger),
		options...,
	)
	r := mux.NewRouter()
	r.Methods("POST").Path(path).Handler(handler)
	return r
}

func makeDecodeRequest(log log.Logger) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		var req entity.Iso8583
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			level.Debug(log).Log("ERROR Decoding request", err)
			return nil, err
		}
		err := validator.New().Struct(req)
		if err != nil {
			return nil, err
		}
		return req, nil
	}
}

func makeEncodeResponse(log log.Logger) httptransport.EncodeResponseFunc {
	return func(c context.Context, w http.ResponseWriter, res interface{}) error {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			level.Debug(log).Log("ERROR Encoding response", err)
			return err
		}

		return nil
	}
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	// switch err {
	// case ErrNotFound:
	// 	return http.StatusNotFound
	// case ErrAlreadyExists, ErrInconsistentIDs:
	// 	return http.StatusBadRequest
	// default:
	// 	return http.StatusInternalServerError
	// }
	return http.StatusInternalServerError
}
