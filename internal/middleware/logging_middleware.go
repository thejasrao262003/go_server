package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // default
		}

		log.Printf("➡ %s %s", r.Method, r.URL.Path)

		next(rw, r)

		log.Printf("⬅ %d %s (%v)",
			rw.statusCode,
			http.StatusText(rw.statusCode),
			time.Since(start),
		)
	}
}
