package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapped, r)
		slog.Info(
			fmt.Sprintf("%d", wrapped.statusCode),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.String("duration", time.Since(start).String()),
		)
	})
}
