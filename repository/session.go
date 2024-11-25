package repository

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"strings"
)

type Session struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitSessionRepo(db *sql.DB, log *zap.Logger) *Session {
	return &Session{Db: db, Logger: log}
}

func (repo *Session) Validate(sessionId string) error {
	isValid := false
	query := `SELECT exists(SELECT 1 FROM sessions WHERE token = $1 AND now() <= expired_at)`
	repo.Logger.Debug("database", zap.String("query", strings.Replace(query, "$1", sessionId, 1)))
	if err := repo.Db.QueryRow(query, sessionId).Scan(&isValid); err != nil {
		return err
	}

	if !isValid {
		return errors.New("invalid session token")
	}
	return nil
}
