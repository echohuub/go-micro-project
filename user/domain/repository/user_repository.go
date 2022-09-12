package repository

import (
	"github.com/heqingbao/go-micro-project/user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	InitTable() error
	FindUserByName(string) (*model.User, error)
	FindUserByID(int64) (*model.User, error)
	CreateUser(user *model.User) (int64, error)
	DeleteUserByID(int64) error
	UpdateUser(user *model.User) error
	FindAll() ([]model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(model.User{}).Error
}

func (u *UserRepository) FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	err := u.mysqlDb.Where("user_name = ?", name).Find(user).Error
	return user, err
}

func (u *UserRepository) FindUserByID(id int64) (*model.User, error) {
	user := &model.User{}
	err := u.mysqlDb.First(user, id).Error
	return user, err
}

func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	err := u.mysqlDb.Create(user).Error
	return user.ID, err
}

func (u *UserRepository) DeleteUserByID(id int64) error {
	return u.mysqlDb.Where("id = ?", id).Delete(model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

func (u *UserRepository) FindAll() ([]model.User, error) {
	users := make([]model.User, 0)
	return users, u.mysqlDb.Find(users).Error
}
