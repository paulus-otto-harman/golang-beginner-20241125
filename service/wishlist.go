package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type WishlistService struct {
	Wishlist *repository.Wishlist
	Logger   *zap.Logger
}

func InitWishlistService(repo repository.Repository, log *zap.Logger) *WishlistService {
	return &WishlistService{Wishlist: repo.Wishlist, Logger: log}
}

func (repo *WishlistService) Store(wishlistItem *domain.WishlistItem, authToken string) error {
	return repo.Wishlist.Store(wishlistItem, authToken)
}

func (repo *WishlistService) Destroy(wishlistItem *domain.WishlistItem, authToken string) error {
	return repo.Wishlist.Destroy(wishlistItem, authToken)
}
