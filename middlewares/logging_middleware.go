package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestTime := time.Now()
		next.ServeHTTP(w, r)

		log.Printf("%s %s %s", requestTime, r.Method, r.URL.Path)
	})
}
