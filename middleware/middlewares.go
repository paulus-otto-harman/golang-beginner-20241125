package middleware

import (
	"go.uber.org/zap"
	"project/service"
)

type Middleware struct {
	AccessLog AccessLog
	Auth      Auth
}

func InitMiddlewares(services service.Service, log *zap.Logger) Middleware {
	return Middleware{
		AccessLog: InitAccessLogMiddleware(log),
		Auth:      InitAuthMiddleware(services, log),
	}

}
