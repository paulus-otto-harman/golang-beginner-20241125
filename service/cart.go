package service

import (
	"go.uber.org/zap"
	"project/domain"
	"project/repository"
)

type CartServiceInterface interface {
	Get(authToken string) (*domain.Cart, error)
	Store(cartItem domain.CartItem, authToken string) error
	Update(cartItem domain.CartItem, authToken string) error
	Delete(productId int, authToken string) error
}

type cartService struct {
	Cart   *repository.Cart
	Logger *zap.Logger
}

func InitCartService(repo repository.Repository, log *zap.Logger) CartServiceInterface {
	return &cartService{Cart: repo.Cart, Logger: log}
}

func (repo *cartService) Get(authToken string) (*domain.Cart, error) {
	cart, err := repo.Cart.Get(authToken)

	if err != nil {
		repo.Logger.Error("get cart error", zap.Error(err))
		return nil, err
	}

	cart.ItemCount = len(cart.Items)
	return cart, nil
}

func (repo *cartService) Store(cartItem domain.CartItem, authToken string) error {
	return repo.Cart.Store(cartItem, authToken)
}

func (repo *cartService) Update(cartItem domain.CartItem, authToken string) error {
	return repo.Cart.Update(cartItem, authToken)
}

func (repo *cartService) Delete(productId int, authToken string) error {
	return repo.Cart.Destroy(productId, authToken)
}
