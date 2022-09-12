package service

import (
	"github.com/heqingbao/go-micro-service-cart/domain/model"
	"github.com/heqingbao/go-micro-service-cart/domain/repository"
)

type ICartService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartService(cartRepository repository.ICartRepository) ICartService {
	return &CartService{CartRepository: cartRepository}
}

type CartService struct {
	CartRepository repository.ICartRepository
}

func (u *CartService) AddCart(product *model.Cart) (id int64, err error) {
	return u.CartRepository.CreateCart(product)
}

func (u *CartService) DeleteCart(id int64) error {
	return u.CartRepository.DeleteCartByID(id)
}

func (u *CartService) UpdateCart(product *model.Cart) error {
	return u.CartRepository.UpdateCart(product)
}

func (u *CartService) FindCartByID(id int64) (*model.Cart, error) {
	return u.CartRepository.FindCartByID(id)
}

func (u *CartService) FindAllCart(userId int64) ([]model.Cart, error) {
	return u.CartRepository.FindAll(userId)
}

func (u *CartService) CleanCart(userId int64) error {
	return u.CartRepository.CleanCart(userId)
}
func (u *CartService) IncrNum(cartId int64, num int64) error {
	return u.CartRepository.IncrNum(cartId, num)
}
func (u *CartService) DecrNum(cartId int64, num int64) error {
	return u.CartRepository.DecrNum(cartId, num)
}
