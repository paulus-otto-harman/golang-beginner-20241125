package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"project/domain"
)

type Address struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitAddressRepo(db *sql.DB, log *zap.Logger) *Address {
	return &Address{Db: db, Logger: log}
}

func (repo *Address) Index(authToken string) ([]domain.Address, error) {
	query := `SELECT id,fullname,email,detail,is_default
				FROM addresses
				JOIN customer_tokens ON addresses.customer_id = customer_tokens.customer_id
				WHERE customer_tokens.token = $1`
	rows, err := repo.Db.Query(query, authToken)
	var addresses []domain.Address
	for rows.Next() {
		var address domain.Address
		if err = rows.Scan(&address.Id, &address.FullName, &address.Email, &address.Detail, &address.IsDefault); err != nil {
			return []domain.Address{}, err
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

func (repo *Address) Store(address *domain.Address, authToken string) error {
	query := `INSERT INTO addresses (customer_id, fullname, email, detail) 
				SELECT customer_id, $1, $2, $3
				FROM customer_tokens
				WHERE token = $4`
	_, err := repo.Db.Exec(query, address.FullName, address.Email, address.Detail, authToken)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Address) Update(address *domain.Address, authToken string) error {
	query := `UPDATE addresses
				SET fullname = $2, email = $3, detail = $4
				FROM customer_tokens
				WHERE addresses.customer_id=customer_tokens.customer_id
				  	AND customer_tokens.token = $5
				    AND addresses.id = $1
			`
	_, err := repo.Db.Exec(query, address.Id, address.FullName, address.Email, address.Detail, authToken)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Address) SetDefault(addressId int, authToken string) error {
	query := `UPDATE addresses
				SET is_default = CASE id WHEN $1 THEN true ELSE false END
				FROM customer_tokens
				WHERE addresses.customer_id = customer_tokens.customer_id
				AND customer_tokens.token = $2`
	_, err := repo.Db.Exec(query, addressId, authToken)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Address) Destroy(addressId int, authToken string) error {
	query := `DELETE FROM addresses
				USING customer_tokens
				WHERE addresses.customer_id = customer_tokens.customer_id
				  AND addresses.id = $1
				  AND customer_tokens.token = $2
				  AND NOT addresses.is_default`
	_, err := repo.Db.Exec(query, addressId, authToken)
	if err != nil {
		return err
	}
	return nil
}
