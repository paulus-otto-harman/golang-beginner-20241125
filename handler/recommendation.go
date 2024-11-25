package handler

import (
	"go.uber.org/zap"
	"net/http"
	"project/service"
	"project/util"
)

type RecommendationHandler struct {
	RecommendationService *service.RecommendationService
	Logger                *zap.Logger
}

func InitRecommendationHandler(service service.Service, log *zap.Logger) RecommendationHandler {
	return RecommendationHandler{RecommendationService: service.Recommendation, Logger: log}
}

func (handle RecommendationHandler) Get(w http.ResponseWriter, r *http.Request) {
	recommendations, _ := handle.RecommendationService.Get()
	util.Response(w).Json(http.StatusOK, "Get all recommended products", recommendations)
}
