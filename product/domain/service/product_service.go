package service

import (
	"github.com/heqingbao/go-micro-project/product/domain/model"
	"github.com/heqingbao/go-micro-project/product/domain/repository"
)

type IProductService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindAllProduct() ([]model.Product, error)
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &ProductService{ProductRepository: productRepository}
}

type ProductService struct {
	ProductRepository repository.IProductRepository
}

func (u *ProductService) AddProduct(product *model.Product) (id int64, err error) {
	return u.ProductRepository.CreateProduct(product)
}

func (u *ProductService) DeleteProduct(id int64) error {
	return u.ProductRepository.DeleteProductByID(id)
}

func (u *ProductService) UpdateProduct(product *model.Product) error {
	return u.ProductRepository.UpdateProduct(product)
}

func (u *ProductService) FindProductByID(id int64) (*model.Product, error) {
	return u.ProductRepository.FindProductByID(id)
}

func (u *ProductService) FindAllProduct() ([]model.Product, error) {
	return u.ProductRepository.FindAll()
}
