package timetransport

import (
	"io"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func loggingMW(logger kitlog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lrw := newLogResponseWriter(w)
			next.ServeHTTP(lrw, r)

			body, _ := io.ReadAll(r.Body)

			if lrw.statusCode > 300 && lrw.statusCode < 600 {
				logger.Log(
					"msg", "failed request",
					"method", r.Method,
					"endpoint", r.URL.Path,
					"statusCode", lrw.statusCode,
					"body", string(body),
				)
			}
		})
	}
}

func newLogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{
		ResponseWriter: w,
		statusCode:     0,
	}
}

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *logResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *logResponseWriter) Write(bytes []byte) (int, error) {
	return lrw.ResponseWriter.Write(bytes)
}
