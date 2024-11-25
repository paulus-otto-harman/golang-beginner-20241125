package repository

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"project/domain"
	"strings"
)

type Auth struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitAuthRepo(db *sql.DB, log *zap.Logger) *Auth {
	return &Auth{Db: db, Logger: log}
}

func (repo *Auth) Authenticate(user domain.User, sessionLifetime int) (*domain.Session, error) {
	query := `INSERT INTO sessions (user_id, expired_at)
    			    SELECT id, now() + make_interval(mins => $3)
    			    FROM users WHERE username=$1 AND password=$2
    			    RETURNING token, expired_at`
	var session domain.Session
	repo.Logger.Debug("database", zap.String("query", strings.Replace(strings.Replace(query, "$1", user.Username, 1), "$2", user.Password, 1)))
	if err := repo.Db.QueryRow(query, user.Username, user.Password, sessionLifetime).Scan(&session.Token, &session.ExpiredAt); err != nil {
		repo.Logger.Error("database query", zap.Error(err))
		return nil, err
	}
	repo.Logger.Info("authentication success", zap.String("username", user.Username))
	return &session, nil
}

func (repo *Auth) Logout(token string) error {
	query := `UPDATE sessions SET expired_at=now() WHERE token=$1`
	repo.Logger.Debug("database", zap.String("query", strings.Replace(query, "$1", token, 1)))
	result, err := repo.Db.Exec(query, token)
	if err != nil {
		repo.Logger.Error("database query", zap.Error(err))
		return err
	}

	if loggedOut, _ := result.RowsAffected(); loggedOut == 0 {
		repo.Logger.Error("database query", zap.Error(errors.New("invalid session")))
		return errors.New("logout failed :: invalid session")
	}

	return nil
}
