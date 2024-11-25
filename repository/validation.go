package repository

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

type Validation struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitValidationRepo(db *sql.DB, log *zap.Logger) *Validation {
	return &Validation{Db: db, Logger: log}
}

func (repo *Validation) IsUnique(tableName string, columnName string, value string) (bool, error) {
	query := `SELECT NOT EXISTS(SELECT 1 FROM` + fmt.Sprintf(" %s ", tableName) + `WHERE ` + columnName + ` = $1)`
	repo.Logger.Debug("database", zap.String("query", strings.Replace(query, "$1", value, 1)))
	var notFound bool
	if err := repo.Db.QueryRow(query, value).Scan(&notFound); err != nil {
		repo.Logger.Error("database", zap.String("query", strings.Replace(query, "$1", value, 1)), zap.Error(err))
		return false, err
	}
	return notFound, nil
}

func (repo *Validation) Exists(tableName string, columnName string, value string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM` + fmt.Sprintf(" %s ", tableName) + `WHERE ` + columnName + ` = $1)`
	repo.Logger.Debug("database", zap.String("query", strings.Replace(query, "$1", value, 1)))
	var found bool
	if err := repo.Db.QueryRow(query, value).Scan(&found); err != nil {
		return false, err
	}
	return found, nil
}

func (repo *Validation) ExistsForUser(authToken string, addressId int) (bool, error) {
	query := `SELECT EXISTS(
					SELECT addresses.customer_id
					FROM addresses
					JOIN (SELECT customer_id AS id FROM customer_tokens WHERE token=$1) customer
					ON addresses.customer_id=customer.id
					WHERE addresses.id=$2
				)`
	var found bool
	if err := repo.Db.QueryRow(query, authToken, addressId).Scan(&found); err != nil {
		return false, err
	}
	return found, nil
}

func (repo *Validation) NotEmptyCart(authToken string) (bool, error) {
	query := `SELECT EXISTS(
				SELECT cart_items.product_id
				FROM cart_items
				JOIN (SELECT id FROM shopping_session_tokens WHERE token=$1) shopping_session
				ON cart_items.shopping_session_id=shopping_session.id)`
	var found bool
	if err := repo.Db.QueryRow(query, authToken).Scan(&found); err != nil {
		return false, err
	}
	return found, nil
}
