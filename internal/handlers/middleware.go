package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrxacker/go-myapp/pkg/logger"
)

// LoggingMiddleware logs HTTP requests with structured fields
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Get request ID from Chi middleware
		requestID := middleware.GetReqID(r.Context())

		// Wrap response writer to capture status code
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		// Process request
		next.ServeHTTP(ww, r)

		// Log with structured fields
		duration := time.Since(start)
		logger.Get().Info("HTTP request",
			logger.String("request_id", requestID),
			logger.String("method", r.Method),
			logger.String("path", r.URL.Path),
			logger.String("remote_addr", r.RemoteAddr),
			logger.Int("status", ww.Status()),
			logger.Int("bytes", ww.BytesWritten()),
			logger.Duration("duration", duration),
			logger.String("user_agent", r.UserAgent()),
		)
	})
}
