package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"project/domain"
)

type Order struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitOrderRepo(db *sql.DB, log *zap.Logger) *Order {
	return &Order{Db: db, Logger: log}
}

func (repo *Order) Create(order *domain.Order) error {
	tx, err := repo.Db.Begin()
	repo.Logger.Debug("database", zap.String("query", "BEGIN"))
	if err != nil {
		repo.Logger.Error("database transaction", zap.Error(err))
		return err
	}

	queryOrder := `INSERT INTO orders (customer_id, address_id, coupon_code, shipping, notes, payment_method)
					SELECT customer_id, $2, $3, $4, $5, $6
					    FROM customer_tokens
						WHERE customer_tokens.token=$1
					RETURNING orders.id
				`

	err = repo.Db.QueryRow(queryOrder, order.CustomerToken, order.AddressId, order.CouponCode, order.Shipping, order.Notes, order.PaymentMethod).Scan(&order.Id)
	if err != nil {
		return err
	}
	repo.Logger.Debug("order created", zap.Int("orderId", order.Id))
	queryOrderItem := `INSERT INTO order_items (order_id, product_id, quantity, unit_price, discount_rate, net_price) 
						SELECT $1,cart_items.product_id, cart_items.quantity, products.price,
						       products.discount_rate, products.price-products.price*products.discount_rate/100
						    FROM cart_items
							JOIN products ON products.id = cart_items.product_id
						    JOIN (SELECT shopping_session_id AS id
						        	FROM shopping_session_tokens
						        	WHERE shopping_session_tokens.token=$2
						    ) shopping_session ON cart_items.shopping_session_id = shopping_session.id
						`
	_, err = repo.Db.Exec(queryOrderItem, order.Id, order.CustomerToken)
	if err != nil {
		repo.Logger.Error("database query", zap.Error(err))
		tx.Rollback()
		repo.Logger.Debug("database", zap.String("query", "ROLLBACK"))
		return err
	}

	queryEmptyCart := `DELETE FROM cart_items 
       					USING shopping_session_tokens
       					WHERE cart_items.shopping_session_id=shopping_session_tokens.shopping_session_id
       					AND shopping_session_tokens.token=$1`
	_, err = repo.Db.Exec(queryEmptyCart, order.CustomerToken)

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
