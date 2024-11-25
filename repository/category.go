package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"project/domain"
)

type Category struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitCategoryRepo(db *sql.DB, log *zap.Logger) *Category {
	return &Category{Db: db, Logger: log}
}

func (repo *Category) Get() ([]domain.Category, error) {
	query := `SELECT id,name FROM categories`
	repo.Logger.Debug("database", zap.String("query", query))
	rows, err := repo.Db.Query(query)
	if err != nil {
		repo.Logger.Error("database query", zap.Error(err))
		return nil, err
	}
	categories := []domain.Category{}
	for rows.Next() {
		var category domain.Category
		if err := rows.Scan(&category.Id, &category.Name); err != nil {
			repo.Logger.Error("database rows scan", zap.Error(err))
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
