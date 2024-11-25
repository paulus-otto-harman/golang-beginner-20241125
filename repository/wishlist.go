package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"log"
	"project/domain"
)

type Wishlist struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitWishlistRepo(db *sql.DB, log *zap.Logger) *Wishlist {
	return &Wishlist{Db: db, Logger: log}
}

func (repo *Wishlist) Store(wishlistItem *domain.WishlistItem, authToken string) error {
	query := `INSERT INTO wishlists (customer_id, product_id)
				SELECT customer_id, $1
				FROM customer_tokens WHERE token = $2
				RETURNING wishlists.created_at`

	if err := repo.Db.QueryRow(query, wishlistItem.ProductId, authToken).Scan(&wishlistItem.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (repo *Wishlist) Destroy(wishlistItem *domain.WishlistItem, authToken string) error {
	query := `DELETE FROM wishlists
				USING customer_tokens
				WHERE wishlists.customer_id = customer_tokens.customer_id 
  				AND wishlists.product_id = $1 AND customer_tokens.token = $2`
	if _, err := repo.Db.Exec(query, wishlistItem.ProductId, authToken); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
