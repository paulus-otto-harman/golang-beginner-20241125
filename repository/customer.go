package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"project/domain"
	"strconv"
	"strings"
)

type Customer struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitCustomerRepo(db *sql.DB, log *zap.Logger) *Customer {
	return &Customer{Db: db, Logger: log}
}

func (repo *Customer) Register(customer *domain.Customer) error {
	tx, err := repo.Db.Begin()
	repo.Logger.Debug("database", zap.String("query", "BEGIN"))
	if err != nil {
		repo.Logger.Error("database transaction", zap.Error(err))
		return err
	}

	query := "INSERT INTO users (username, password) VALUES($1, $2) RETURNING id, created_at, ''"
	repo.Logger.Debug("database", zap.String("query", strings.Replace(strings.Replace(query, "$1", customer.Username, 1), "$2", customer.Password, 1)))
	var userId int
	err = repo.Db.QueryRow(query, customer.Username, customer.Password).Scan(&userId, &customer.CreatedAt, &customer.Password)
	if err != nil {
		repo.Logger.Error("database query", zap.Error(err))
		return err
	}

	query = "INSERT INTO customers (user_id, name) VALUES($1, $2)"
	repo.Logger.Debug("database", zap.String("query", strings.Replace(strings.Replace(query, "$1", strconv.Itoa(userId), 1), "$2", customer.Name, 1)))
	_, err = repo.Db.Exec(query, userId, customer.Name)
	if err != nil {
		repo.Logger.Error("database query", zap.Error(err))
		tx.Rollback()
		repo.Logger.Debug("database", zap.String("query", "ROLLBACK"))
		return err
	}

	if err = tx.Commit(); err != nil {
		repo.Logger.Error("commit", zap.Error(err))
	}

	repo.Logger.Debug("database", zap.String("query", "COMMIT"))
	return nil
}

func (repo *Customer) Profile(authToken string) (*domain.Customer, error) {
	query := `SELECT customers.name,users.username,users.created_at FROM customers 
				JOIN users ON users.id = customers.user_id
				JOIN (SELECT user_id
					FROM sessions
					WHERE sessions.token=$1
				) session ON users.id = session.user_id`
	var customer domain.Customer
	if err := repo.Db.QueryRow(query, authToken).Scan(&customer.Name, &customer.Username, &customer.CreatedAt); err != nil {
		repo.Logger.Error("database query", zap.Error(err))
		return nil, err
	}
	return &customer, nil
}
