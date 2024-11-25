package service

import (
	"go.uber.org/zap"
	"project/config"
	"project/repository"
)

type Service struct {
	Address        AddressServiceInterface
	Auth           AuthServiceInterface
	Banner         *BannerService
	Bestseller     *BestsellerService
	Cart           *CartService
	Category       *CategoryService
	Order          *OrderService
	Customer       *CustomerService
	Product        *ProductService
	Recommendation *RecommendationService
	Session        *SessionService
	Validation     *ValidationService
	Wishlist       *WishlistService
	Weekly         *WeeklyService
}

// TODO
func InitServices(repositories repository.Repository, log *zap.Logger, config config.AppConfig) Service {
	return Service{
		Address:        InitAddressService(repositories, log),
		Auth:           InitAuthService(repositories, log, config),
		Banner:         InitBannerService(repositories, log),
		Bestseller:     InitBestsellerService(repositories, log),
		Cart:           InitCartService(repositories, log),
		Category:       InitCategoryService(repositories, log),
		Order:          InitOrderService(repositories, log),
		Customer:       InitCustomerService(repositories, log),
		Product:        InitProductService(repositories, log),
		Recommendation: InitRecommendationService(repositories, log),
		Session:        InitSessionService(repositories, log),
		Validation:     InitValidationService(repositories, log),
		Wishlist:       InitWishlistService(repositories, log),
		Weekly:         InitWeeklyService(repositories, log),
	}
}
