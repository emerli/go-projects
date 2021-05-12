package middlewares

import (
	"log"
	"net/http"
)

func LogAuditMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s@%s] %s\n", r.Header.Get("username"), r.Host, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
