package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"project/domain"
)

type Banner struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitBannerRepo(db *sql.DB, log *zap.Logger) *Banner {
	return &Banner{Db: db, Logger: log}
}

func (repo *Banner) Get() ([]domain.Banner, error) {
	query := `SELECT photo, title, subtitle, path_page FROM banners WHERE started_at <= now() AND finished_at >= now()`
	repo.Logger.Debug("database", zap.String("query", query))
	rows, err := repo.Db.Query(query)
	if err != nil {
		return nil, err
	}
	banners := []domain.Banner{}
	for rows.Next() {
		var banner domain.Banner
		if err := rows.Scan(&banner.Photo, &banner.Title, &banner.Subtitle, &banner.PathPage); err != nil {
			return nil, err
		}
		banners = append(banners, banner)
	}
	return banners, nil
}
