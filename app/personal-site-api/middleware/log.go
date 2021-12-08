package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
		// TODO: better logging
		fmt.Printf("%s: %s %dms\n", r.Method, r.RequestURI, time.Now().UnixMilli()-start.UnixMilli())

	})
}
