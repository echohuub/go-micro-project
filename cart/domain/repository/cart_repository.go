package repository

import (
	"errors"
	"github.com/heqingbao/go-micro-service-cart/domain/model"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCart(*model.Cart) (int64, error)
	DeleteCartByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

func (u *CartRepository) InitTable() error {
	return u.mysqlDb.CreateTable(model.Cart{}).Error
}

func (u *CartRepository) FindCartByID(id int64) (*model.Cart, error) {
	Cart := &model.Cart{}
	err := u.mysqlDb.First(Cart, id).Error
	return Cart, err
}

func (u *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := u.mysqlDb.FirstOrCreate(cart, model.Cart{
		ProductID: cart.ProductID,
		SizeID:    cart.SizeID,
		UserID:    cart.UserID,
	})
	if db.Error != nil {
		return 0, db.Error
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("购物车插入失败")
	}
	return cart.ID, nil
}

func (u *CartRepository) DeleteCartByID(id int64) error {
	return u.mysqlDb.Where("id = ?", id).Delete(&model.Cart{}).Error
}

func (u *CartRepository) UpdateCart(Cart *model.Cart) error {
	return u.mysqlDb.Model(Cart).Update(&Cart).Error
}

func (u *CartRepository) FindAll(userId int64) ([]model.Cart, error) {
	categories := make([]model.Cart, 0)
	return categories, u.mysqlDb.Where("user_id = ?", userId).
		Find(categories).Error
}

func (u *CartRepository) CleanCart(userId int64) error {
	return u.mysqlDb.Where("user_id = ?", userId).Delete(&model.Cart{}).Error
}

func (u *CartRepository) IncrNum(cartId int64, num int64) error {
	cart := &model.Cart{ID: cartId}
	return u.mysqlDb.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (u *CartRepository) DecrNum(cartId int64, num int64) error {
	cart := &model.Cart{ID: cartId}
	db := u.mysqlDb.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}
