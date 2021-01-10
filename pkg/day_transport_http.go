package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-kit/kit/transport"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s Service, logger kitlog.Logger) http.Handler {
	// TODO: add auth MW (https://github.com/go-kit/kit/tree/master/auth/jwt)
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

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
	r.Use(contentTypeMW)

	// routes
	dayRouter := r.Path("/v1/lists/{listId}/days").Subrouter()
	dayRouter.Handle("/{date}", addDayHandler).Methods(http.MethodPut)
	dayRouter.Handle("", getDaysHandler).Methods(http.MethodGet)
	dayRouter.Handle("/{date}", getDayHandler).Methods(http.MethodGet)
	dayRouter.Handle("/{date}", updateDayHandler).Methods(http.MethodPatch)
	dayRouter.Handle("/{date}", deleteDayHandler).Methods(http.MethodDelete)

	return r
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

var errBadRoute = errors.New("bad route")

// days
const dateLayout = "2006-02-01"

//AddDay(ctx context.Context, listID string, date time.Time, moments []Moment) (Day, error)
func decodeAddDayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["listId"]
	if !ok {
		return nil, errBadRoute
	}

	dateString, ok := vars["date"]
	if !ok {
		return nil, errBadRoute
	}

	date, err := time.Parse(dateLayout, dateString)
	if err != nil {
		return nil, err
	}

	var body struct {
		Moments []Moment `json:"moments"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return addDayRequest{
		ListID:  id,
		Date:    date,
		Moments: body.Moments,
	}, nil
}

//GetDays(ctx context.Context, listID string, opts ...ListDaysOption) ([]Day, error)
func decodeGetDaysRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, ok := mux.Vars(r)["listId"]
	if !ok {
		return nil, errBadRoute
	}

	return getDaysRequest{
		ListID: id,
	}, nil
}

//GetDay(ctx context.Context, listID string, date time.Time) (Day, error)
func decodeGetDayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["listId"]
	if !ok {
		return nil, errBadRoute
	}

	dateString, ok := vars["date"]
	if !ok {
		return nil, errBadRoute
	}

	date, err := time.Parse(dateLayout, dateString)
	if err != nil {
		return nil, err
	}

	return getDayRequest{
		ListID: id,
		Date:   date,
	}, nil
}

//UpdateDay(ctx context.Context, listID string, date time.Time, moments []Moment) (Day, error)
func decodeUpdateDayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	id, ok := vars["listId"]
	if !ok {
		return nil, errBadRoute
	}

	dateString, ok := vars["date"]
	if !ok {
		return nil, errBadRoute
	}

	date, err := time.Parse(dateLayout, dateString)
	if err != nil {
		return nil, err
	}

	var body struct {
		Moments []Moment `json:"moments"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return updateDayRequest{
		ListID:  id,
		Date:    date,
		Moments: body.Moments,
	}, nil
}

//DeleteDay(ctx context.Context, listID string, date time.Time) error
func decodeDeleteDayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["listId"]
	if !ok {
		return nil, errBadRoute
	}

	dateString, ok := vars["date"]
	if !ok {
		return nil, errBadRoute
	}

	date, err := time.Parse(dateLayout, dateString)
	if err != nil {
		return nil, err
	}

	return deleteDayRequest{
		ListID: id,
		Date:   date,
	}, nil
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
