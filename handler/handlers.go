package handler

import (
	"go.uber.org/zap"
	"project/service"
	"project/util"
)

type Handler struct {
	Address        AddressHandler
	Auth           AuthHandler
	Banner         BannerHandler
	Bestseller     BestsellerHandler
	Category       CategoryHandler
	Cart           CartHandler
	Order          OrderHandler
	Customer       CustomerHandler
	Product        ProductHandler
	Recommendation RecommendationHandler
	Wishlist       WishlistHandler
	Weekly         WeeklyHandler
}

func InitHandlers(services service.Service, log *zap.Logger, validator util.Validation) Handler {
	return Handler{
		Address:        InitAddressHandler(services, log, validator),
		Auth:           InitAuthHandler(services, log, validator),
		Banner:         InitBannerHandler(services, log),
		Bestseller:     InitBestsellerHandler(services, log),
		Cart:           InitCartHandler(services, log),
		Order:          InitOrderHandler(services, log, validator),
		Category:       InitCategoryHandler(services, log),
		Customer:       InitCustomerHandler(services, log, validator),
		Product:        InitProductHandler(services, log),
		Recommendation: InitRecommendationHandler(services, log),
		Wishlist:       InitWishlistHandler(services, log, validator),
		Weekly:         InitWeeklyHandler(services, log),
	}

}
