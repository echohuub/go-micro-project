package repository

import (
	"github.com/heqingbao/go-micro-project/product/domain/model"
	"github.com/jinzhu/gorm"
)

type IProductRepository interface {
	InitTable() error
	FindProductByID(int64) (*model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	DeleteProductByID(int64) error
	UpdateProduct(*model.Product) error
	FindAll() ([]model.Product, error)
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

func (u *ProductRepository) InitTable() error {
	return u.mysqlDb.CreateTable(
		model.Product{},
		model.ProductImage{},
		model.ProductSize{},
		model.ProductSeo{}).Error
}

func (u *ProductRepository) FindProductByID(id int64) (*model.Product, error) {
	Product := &model.Product{}
	err := u.mysqlDb.Preload("Image").
		Preload("Size").
		Preload("Seo").
		First(Product, id).Error
	return Product, err
}

func (u *ProductRepository) CreateProduct(Product *model.Product) (int64, error) {
	err := u.mysqlDb.Create(Product).Error
	return Product.ID, err
}

func (u *ProductRepository) DeleteProductByID(id int64) error {
	// 开启事务
	tx := u.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	// 删除
	if err := tx.Unscoped().Where("id = ?", id).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("product_id = ?", id).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("product_id = ?", id).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("product_id = ?", id).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (u *ProductRepository) UpdateProduct(Product *model.Product) error {
	return u.mysqlDb.Model(Product).Update(&Product).Error
}

func (u *ProductRepository) FindAll() ([]model.Product, error) {
	categories := make([]model.Product, 0)
	return categories, u.mysqlDb.Preload("image").
		Preload("size").
		Preload("seo").
		Find(categories).Error
}
