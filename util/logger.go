package util

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		handler.ServeHTTP(w, r)
		duration := time.Since(startTime)
		// log.Print("Test")

		logger := log.Info()
		logger.Str("protocol", "HTTP").
			Str("method", r.Method).
			// Int("status", 200).
			Str("path", r.RequestURI).
			// Str("query", r.URL.RawQuery).
			// Str("ip", r.RemoteAddr).
			// Str("user_agent", r.UserAgent()).
			// Str("status_text", http.StatusText(200)).
			Dur("duration", duration).
			Msg("received a HTTP request")
	})
}
