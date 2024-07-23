package middlewares

import (
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(rw, r)

		switch rw.statusCode {
		case http.StatusBadRequest:
			http.Error(w, `{"error": "Bad Request"}`, http.StatusBadRequest)
		case http.StatusUnauthorized:
			http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		case http.StatusForbidden:
			http.Error(w, `{"error": "Forbidden"}`, http.StatusForbidden)
		case http.StatusNotFound:
			http.Error(w, `{"error": "Not Found"}`, http.StatusNotFound)
		case http.StatusInternalServerError:
			http.Error(w, `{"error": "Internal Server Error"}`, http.StatusInternalServerError)
		default:
		}
	})
}
