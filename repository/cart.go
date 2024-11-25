package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"log"
	"project/domain"
	"strconv"
)

type Cart struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitCartRepo(db *sql.DB, log *zap.Logger) *Cart {
	return &Cart{Db: db, Logger: log}
}

func (repo *Cart) Get(authToken string) (*domain.Cart, error) {
	query := `SELECT product_id, quantity
				FROM cart_items
				JOIN shopping_session_tokens ON cart_items.shopping_session_id=shopping_session_tokens.shopping_session_id
				WHERE shopping_session_tokens.token = $1`
	repo.Logger.Debug("database", zap.String("query", query))
	rows, err := repo.Db.Query(query, authToken)

	if err != nil {
		repo.Logger.Error("query failed", zap.Error(err))
		return nil, err
	}
	cart := domain.Cart{}
	for rows.Next() {
		var cartItem domain.CartItem
		if err = rows.Scan(&cartItem.ProductId, &cartItem.Quantity); err != nil {
			repo.Logger.Error("row scan failed", zap.Error(err))
			return nil, err
		}
		cart.Items = append(cart.Items, cartItem)
	}

	return &cart, nil
}

func (repo *Cart) Store(cartItem domain.CartItem, authToken string) error {
	queryInsertShoppingSession := `INSERT INTO shopping_sessions (customer_id)
										SELECT customer_id FROM customer_tokens
										WHERE customer_tokens.token=$3
										ON CONFLICT (customer_id) DO NOTHING
									RETURNING id`
	querySelectShoppingSession := `SELECT id AS shopping_session_id FROM create_shopping_session
									UNION
									SELECT shopping_session_id FROM shopping_session_tokens WHERE token=$3`
	queryShoppingSession := `WITH create_shopping_session AS (` + queryInsertShoppingSession + `),
								first_or_create_shopping_session AS (` + querySelectShoppingSession + `)`

	quantityNew := strconv.Itoa(cartItem.Quantity)
	quantityUpdate := strconv.Itoa(cartItem.Quantity)
	if cartItem.Quantity == 0 {
		quantityNew = " 1"
		quantityUpdate = "cart_items.quantity + 1"
	}

	queryCartItem := queryShoppingSession + `
				INSERT INTO cart_items (shopping_session_id, product_id, quantity)
				SELECT shopping_session_id, $1,$2
				FROM first_or_create_shopping_session
				ON CONFLICT(shopping_session_id, product_id)
				DO UPDATE SET quantity =` + quantityUpdate

	if _, err := repo.Db.Exec(queryCartItem, cartItem.ProductId, quantityNew, authToken); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (repo *Cart) Update(cartItem domain.CartItem, authToken string) error {
	query := `UPDATE cart_items
				SET quantity = $1
				FROM shopping_session_tokens
				WHERE shopping_session_tokens.shopping_session_id = cart_items.shopping_session_id
				AND product_id = $2
				AND shopping_session_tokens.token = $3`

	if _, err := repo.Db.Exec(query, cartItem.Quantity, cartItem.ProductId, authToken); err != nil {
		return err
	}
	return nil
}

func (repo *Cart) Destroy(productId int, authToken string) error {
	query := `DELETE FROM cart_items
				USING shopping_session_tokens
				WHERE shopping_session_tokens.shopping_session_id = cart_items.shopping_session_id
				AND product_id = $1
				AND shopping_session_tokens.token = $2`

	if _, err := repo.Db.Exec(query, productId, authToken); err != nil {
		return err
	}
	return nil
}
