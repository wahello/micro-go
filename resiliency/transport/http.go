package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"micro-go/resiliency/endpoint"
	"net/http"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

// make http handler use mux
func MakeHttpHandler(ctx context.Context, endpoints endpoint.UseStringEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/op/{type}/{a}/{b}").Handler(kithttp.NewServer(
		endpoints.UseStringEndpoint, decodeStringRequest, encodeStringResponse, options...))

	r.Path("/metrics").Handler(promhttp.Handler())

	// create health check handler
	r.Methods("GET").Path("/health").Handler(kithttp.NewServer(
		endpoints.HealthCheckEndpoint, decodeHealthCheckRequest, encodeStringResponse, option...))

	// 添加hytrix 监控数据
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	r.Handle("/hystrix/stream", hystrixStreamHandler)

	return r
}

func decodeHealthCheckRequest(i context.Context, request2 *http.Request) (request interface{}, err error) {

}

func decodeStringRequest(i context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrorBadRequest
	}

	pa, ok := vars["a"]
	if !ok {

	}
}

func encodeStringResponse(i context.Context, writer http.ResponseWriter, i2 interface{}) error {

}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}
