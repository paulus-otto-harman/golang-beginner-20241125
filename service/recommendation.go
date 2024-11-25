package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type RecommendationService struct {
	Recommendation *repository.Recommendation
	Logger         *zap.Logger
}

func InitRecommendationService(repo repository.Repository, log *zap.Logger) *RecommendationService {
	return &RecommendationService{Recommendation: repo.Recommendation, Logger: log}
}

func (repo *RecommendationService) Get() ([]domain.Recommendation, error) {
	return repo.Recommendation.Get()
}
