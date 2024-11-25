package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"log"
	"project/domain"
)

type Weekly struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitWeeklyRepo(db *sql.DB, log *zap.Logger) *Weekly {
	return &Weekly{Db: db, Logger: log}
}

func (repo *Weekly) Get() ([]domain.Weekly, error) {
	query := `SELECT photo, title, subtitle, product_id
				FROM weeklies
				WHERE started_at <= now() AND finished_at >= now()`
	repo.Logger.Debug("database", zap.String("query", query))
	rows, err := repo.Db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var weeklies []domain.Weekly
	for rows.Next() {
		var weekly domain.Weekly
		if err = rows.Scan(&weekly.Photo, &weekly.Title, &weekly.Subtitle, &weekly.ProductId); err != nil {
			log.Println(err)
			return nil, err
		}
		weeklies = append(weeklies, weekly)
	}
	return weeklies, nil
}
