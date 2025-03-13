package middleware

import (
	"net/http"
	"time"

	"github.com/vishwaszadte/numinaut-be/pkg/logger"
	"go.uber.org/zap"
)

// The LoggingMiddleware function is a middleware function in Go that takes an http.Handler as input.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		logger.Info("Request started",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
		)

		next.ServeHTTP(w, r)

		endTime := time.Now()
		duration := endTime.Sub(startTime)
		logger.Info("Request completed",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.Duration("duration", duration),
		)
	})
}
