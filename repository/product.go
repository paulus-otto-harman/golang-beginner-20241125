package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"math"
	"project/domain"
	"strings"
)

type Product struct {
	Db     *sql.DB
	Logger *zap.Logger
}

func InitProductRepo(db *sql.DB, log *zap.Logger) *Product {
	return &Product{Db: db, Logger: log}
}

func (repo *Product) Index(page int, limit int, keyword string, categoryId int, authToken string) (int, int, []domain.Product, error) {
	whereHasKeyword, orderBy := keywordBuilder(keyword)
	whereHasCategory := categoryBuilder(categoryId)
	sessionToken := tokenBuilder(authToken)

	var count int
	queryCount := `SELECT COUNT(*) FROM products WHERE products.name ILIKE ` + whereHasKeyword + whereHasCategory
	repo.Logger.Debug(queryCount)
	err := repo.Db.QueryRow(queryCount, categoryId).Scan(&count)
	if err != nil {
		repo.Logger.Error(err.Error())
		return 0, 0, nil, err
	}

	query := `SELECT id,name, thumbnail, category_id, rating, price, discount_rate, price-(price*discount_rate/100),
				w.product_id IS NOT NULL, created_at, now()::date-created_at::date<=30
				FROM products
				LEFT JOIN average_product_reviews ON average_product_reviews.product_id = products.id
				LEFT JOIN (SELECT product_id FROM wishlist_tokens WHERE wishlist_tokens.token=$4) w 
				ON products.id=w.product_id
				WHERE name ILIKE ` + whereHasKeyword + whereHasCategory + orderBy + ` LIMIT $3 OFFSET $2`
	offset := (page - 1) * limit

	repo.Logger.Debug("database", zap.String("query", query))
	rows, err := repo.Db.Query(query, categoryId, offset, limit, sessionToken)
	if err != nil {
		return 0, 0, nil, err
	}

	products := []domain.Product{}
	for rows.Next() {
		var product domain.Product
		if err = rows.Scan(&product.Id, &product.Name, &product.ThumbnailUrl, &product.CategoryId, &product.Rating, &product.Price, &product.DiscountRate, &product.NetPrice, &product.OnMyWishlist, &product.CreatedAt, &product.IsNew); err != nil {
			return 0, 0, nil, err
		}

		products = append(products, product)
	}
	return count, int(math.Ceil(float64(count / limit))), products, nil
}

func (repo *Product) Get(product *domain.ProductDetail, authToken string) error {
	sessionToken := tokenBuilder(authToken)

	query := `WITH photos AS (
					SELECT product_photos.product_id, product_photos.photo_url
					FROM product_photos
					ORDER BY product_photos.id
				), variants AS (
				    SELECT product_variants.product_id, product_variants.name AS variant_name
					FROM product_variants
					ORDER BY product_variants.id
				)
				SELECT id,name, thumbnail, category_id, rating, price, discount_rate, price-(price*discount_rate/100),
				w.product_id IS NOT NULL, created_at, now()::date-created_at::date<=30,
				(SELECT string_agg(photo_url, ',') FROM photos WHERE photos.product_id=products.id),
				(SELECT string_agg(variant_name, ',') FROM variants WHERE variants.product_id=products.id)
				FROM products
				LEFT JOIN average_product_reviews ON average_product_reviews.product_id = products.id
				LEFT JOIN (SELECT product_id FROM wishlist_tokens WHERE wishlist_tokens.token=$2) w 
				ON products.id=w.product_id WHERE id=$1`
	repo.Logger.Debug("database", zap.String("query", query))

	var photos sql.NullString
	var variants sql.NullString
	if err := repo.Db.QueryRow(query, product.Id, sessionToken).Scan(&product.Id, &product.Name,
		&product.ThumbnailUrl, &product.CategoryId, &product.Rating, &product.Price, &product.DiscountRate,
		&product.NetPrice, &product.OnMyWishlist, &product.CreatedAt, &product.IsNew,
		&photos, &variants); err != nil {
		return err
	}

	if photos.Valid {
		pics := strings.Split(photos.String, ",")
		for _, pic := range pics {
			product.Photos = append(product.Photos, domain.ProductPhoto{PhotoUrl: pic})
		}
	}

	if variants.Valid {
		productVariants := strings.Split(variants.String, ",")
		for _, variant := range productVariants {
			product.Variants = append(product.Variants, domain.ProductVariant{Name: variant})
		}
	}

	return nil
}

func keywordBuilder(keyword string) (string, string) {
	if keyword == "" {
		return "'%%'", ""
	}
	return "'%" + keyword + "%'", " ORDER BY name"
}

func categoryBuilder(categoryId int) string {
	if categoryId != 0 {
		return ` AND category_id = $1`
	}
	return ` AND category_id != $1`
}

func tokenBuilder(authToken string) string {
	if authToken == "" {
		return "00000000-0000-0000-0000-000000000000"
	}
	return authToken
}
