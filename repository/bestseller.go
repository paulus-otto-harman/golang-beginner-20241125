package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"math"
	"project/domain"
)

type Bestseller struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitBestsellerRepo(db *sql.DB, log *zap.Logger) *Bestseller {
	return &Bestseller{Db: db, Logger: log}
}

func (repo *Bestseller) Index(page int, limit int) (int, int, []domain.Product, error) {
	var count int
	queryCount := `SELECT COUNT(*) FROM products WHERE products.id IN 
					(SELECT DISTINCT product_id
					FROM order_items
					JOIN orders ON order_items.order_id=orders.id
					WHERE now()-orders.created_at <= interval '1 month')`
	err := repo.Db.QueryRow(queryCount).Scan(&count)

	query := `SELECT id,name,category_id FROM products WHERE products.id IN
				 (SELECT DISTINCT product_id
					FROM order_items
					JOIN orders ON order_items.order_id=orders.id
					WHERE now()-orders.created_at <= interval '1 month')
				   LIMIT $2 OFFSET $1`
	repo.Logger.Debug("database", zap.String("query", query))
	offset := (page - 1) * limit
	rows, err := repo.Db.Query(query, offset, limit)
	if err != nil {
		return 0, 0, nil, err
	}

	products := []domain.Product{}
	for rows.Next() {
		var product domain.Product
		if err = rows.Scan(&product.Id, &product.Name, &product.CategoryId); err != nil {
			return 0, 0, nil, err
		}
		products = append(products, product)
	}
	return count, int(math.Ceil(float64(count / limit))), products, nil
}
