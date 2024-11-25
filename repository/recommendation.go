package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"log"
	"project/domain"
)

type Recommendation struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitRecommendationRepo(db *sql.DB, log *zap.Logger) *Recommendation {
	return &Recommendation{Db: db, Logger: log}
}

func (repo *Recommendation) Get() ([]domain.Recommendation, error) {
	query := `SELECT title, subtitle, photo, product_id FROM recommendations`
	repo.Logger.Debug("database", zap.String("query", query))
	rows, err := repo.Db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	recommendations := []domain.Recommendation{}
	for rows.Next() {
		var recommendation domain.Recommendation
		if err = rows.Scan(&recommendation.Title, &recommendation.Subtitle, &recommendation.Photo, &recommendation.ProductId); err != nil {
			log.Println(err)
			return nil, err
		}
		recommendations = append(recommendations, recommendation)
	}
	return recommendations, nil
}
