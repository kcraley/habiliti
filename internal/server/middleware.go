package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// setHeaders is middleware which configures the serve mux with the
// required headers as outlined for the Terraform Registry disovery
// protocol.  This also includes additional configuration for best practices.
func setHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}

// logRequests is a middleware that is used to inject and standardize
// logging for each request.
func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"Content-Type":    r.Header.Get("Content-Type"),
			"Endpoint":        r.URL.Path,
			"Host":            r.Header.Get("Host"),
			"Params":          r.URL.RawQuery,
			"User-Agent":      r.Header.Get("User-Agent"),
			"X-Forwarded-For": r.Header.Get("X-Forwarded-For"),
		}).Infof("received request")
		next.ServeHTTP(w, r)
	})
}
