package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"project/service"
	"project/util"
)

type Auth struct {
	SessionService *service.SessionService
	Logger         *zap.Logger
}

func InitAuthMiddleware(service service.Service, log *zap.Logger) Auth {
	return Auth{SessionService: service.Session, Logger: log}
}

func (middleware *Auth) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("token")
		middleware.Logger.Debug("validate session", zap.String("token", sessionId))

		if err := middleware.SessionService.Validate(sessionId); err != nil {
			util.Response(w).Json(http.StatusUnauthorized, "Unauthorized")
			return
		}
		next.ServeHTTP(w, r)
	})
}
