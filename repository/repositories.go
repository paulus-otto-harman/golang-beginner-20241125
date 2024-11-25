package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"project/config"
)

type Repository struct {
	Address        *Address
	Auth           *Auth
	Banner         *Banner
	Bestseller     *Bestseller
	Cart           *Cart
	Category       *Category
	Order          *Order
	Customer       *Customer
	Product        *Product
	Recommendation *Recommendation
	Session        *Session
	Validation     *Validation
	Wishlist       *Wishlist
	Weekly         *Weekly
}

// TODO
func InitRepositories(db *sql.DB, log *zap.Logger, config config.AppConfig) (Repository, *zap.Logger, config.AppConfig) {
	return Repository{
		Address:        InitAddressRepo(db, log),
		Auth:           InitAuthRepo(db, log),
		Banner:         InitBannerRepo(db, log),
		Bestseller:     InitBestsellerRepo(db, log),
		Cart:           InitCartRepo(db, log),
		Category:       InitCategoryRepo(db, log),
		Order:          InitOrderRepo(db, log),
		Customer:       InitCustomerRepo(db, log),
		Product:        InitProductRepo(db, log),
		Recommendation: InitRecommendationRepo(db, log),
		Session:        InitSessionRepo(db, log),
		Validation:     InitValidationRepo(db, log),
		Wishlist:       InitWishlistRepo(db, log),
		Weekly:         InitWeeklyRepo(db, log),
	}, log, config
}
