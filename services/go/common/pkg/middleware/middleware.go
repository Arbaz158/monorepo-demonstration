package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging wraps a handler to record basic request metadata.
func Logging(l *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		l.Printf("%s %s completed in %s", r.Method, r.URL.Path, time.Since(start))
	})
}
