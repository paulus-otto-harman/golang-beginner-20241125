package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type AccessLog struct {
	Log *zap.Logger
}

func InitAccessLogMiddleware(log *zap.Logger) AccessLog {
	return AccessLog{Log: log}
}

func (middleware *AccessLog) Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		handler.ServeHTTP(w, r)

		duration := time.Since(start)

		middleware.Log.Info("http request", zap.String("method", r.Method), zap.String("url", r.URL.String()), zap.Duration("duration", duration))
	})

}
