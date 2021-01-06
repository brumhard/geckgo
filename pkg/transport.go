package pkg

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"

	"github.com/gorilla/mux"
)

func MakeHandler(s Service, logger log.Logger) http.Handler {
	// call NewServer functions
	// new server for each Handlerfunc for each Endpoint (.../kit/transport/http.NewServer(endpoint,...)

	r := mux.NewRouter()
	r.Use(authWM, contentTypeMW)

	// routes

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

func authWM(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO
		next.ServeHTTP(w, r)
	})
}

func encodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// TODO: decoder functions (take vars, take json body)

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// switch possible errors from business logic
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
