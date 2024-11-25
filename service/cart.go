package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type CartService struct {
	Cart   *repository.Cart
	Logger *zap.Logger
}

func InitCartService(repo repository.Repository, log *zap.Logger) *CartService {
	return &CartService{Cart: repo.Cart, Logger: log}
}

func (repo *CartService) Get(authToken string) (*domain.Cart, error) {
	cart, err := repo.Cart.Get(authToken)

	if err != nil {
		repo.Logger.Error("get cart error", zap.Error(err))
		return nil, err
	}

	cart.ItemCount = len(cart.Items)
	return cart, nil
}

func (repo *CartService) Store(cartItem domain.CartItem, authToken string) error {
	return repo.Cart.Store(cartItem, authToken)
}

func (repo *CartService) Update(cartItem domain.CartItem, authToken string) error {
	return repo.Cart.Update(cartItem, authToken)
}

func (repo *CartService) Delete(productId int, authToken string) error {
	return repo.Cart.Destroy(productId, authToken)
}
