package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var errBadRoute = errors.New("bad route")

func MakeHandler(s Service, logger kitlog.Logger) http.Handler {
	// TODO: add auth MW (https://github.com/go-kit/kit/tree/master/auth/jwt)
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	// list
	addListHandler := kithttp.NewServer(
		makeAddListEndpoint(s),
		decodeAddListRequest,
		encodeJSONResponse,
		opts...,
	)
	getListsHandler := kithttp.NewServer(
		makeGetListsEndpoint(s),
		decodeGetListsRequest,
		encodeJSONResponse,
		opts...,
	)
	getListHandler := kithttp.NewServer(
		makeGetListEndpoint(s),
		decodeGetListRequest,
		encodeJSONResponse,
		opts...,
	)
	updateListHandler := kithttp.NewServer(
		makeUpdateListEndpoint(s),
		decodeUpdateListRequest,
		encodeJSONResponse,
		opts...,
	)
	deleteListHandler := kithttp.NewServer(
		makeDeleteListEndpoint(s),
		decodeDeleteListRequest,
		encodeJSONResponse,
		opts...,
	)

	// day
	addDayHandler := kithttp.NewServer(
		makeAddDayEndpoint(s),
		decodeAddDayRequest,
		encodeJSONResponse,
		opts...,
	)
	getDaysHandler := kithttp.NewServer(
		makeGetDaysEndpoint(s),
		decodeGetDaysRequest,
		encodeJSONResponse,
		opts...,
	)
	getDayHandler := kithttp.NewServer(
		makeGetDayEndpoint(s),
		decodeGetDayRequest,
		encodeJSONResponse,
		opts...,
	)
	updateDayHandler := kithttp.NewServer(
		makeUpdateDayEndpoint(s),
		decodeUpdateDayRequest,
		encodeJSONResponse,
		opts...,
	)
	deleteDayHandler := kithttp.NewServer(
		makeDeleteDayEndpoint(s),
		decodeDeleteDayRequest,
		encodeJSONResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Use(loggingMW(logger), contentTypeMW)

	// routes
	v1Router := r.Path("/v1").Subrouter()

	// list
	v1ListRouter := v1Router.Path("/lists").Subrouter()
	v1ListRouter.Handle("", addListHandler).Methods(http.MethodPost)
	v1ListRouter.Handle("", getListsHandler).Methods(http.MethodGet)
	v1ListRouter.Handle("/{listId}", getListHandler).Methods(http.MethodGet)
	v1ListRouter.Handle("/{listId}", updateListHandler).Methods(http.MethodPatch)
	v1ListRouter.Handle("/{listId}", deleteListHandler).Methods(http.MethodDelete)

	// day
	v1DayRouter := v1Router.Path("/lists/{listId}/days").Subrouter()
	v1DayRouter.Handle("/{date}", addDayHandler).Methods(http.MethodPut)
	v1DayRouter.Handle("", getDaysHandler).Methods(http.MethodGet)
	v1DayRouter.Handle("/{date}", getDayHandler).Methods(http.MethodGet)
	v1DayRouter.Handle("/{date}", updateDayHandler).Methods(http.MethodPatch)
	v1DayRouter.Handle("/{date}", deleteDayHandler).Methods(http.MethodDelete)

	return cutTrailingSlashMW(r)
}

func cutTrailingSlashMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}

func contentTypeMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "unsupported content-type", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func encodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(interface{ error() error }); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

// encodeError is not only called for errors in an endpoint's but also for errors during decoding.
// That's why headers need to be set here as well as in the encodeJSONResponse function.
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// TODO analyze error types from business-logic
	w.WriteHeader(http.StatusInternalServerError)

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
