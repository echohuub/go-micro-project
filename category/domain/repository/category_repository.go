package repository

import (
	"github.com/heqingbao/go-micro-project/category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	InitTable() error
	FindCategoryByID(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryByID(int64) error
	UpdateCategory(*model.Category) error
	FindAll() ([]model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

func (u *CategoryRepository) InitTable() error {
	return u.mysqlDb.CreateTable(model.Category{}).Error
}

func (u *CategoryRepository) FindCategoryByID(id int64) (*model.Category, error) {
	category := &model.Category{}
	err := u.mysqlDb.First(category, id).Error
	return category, err
}

func (u *CategoryRepository) CreateCategory(category *model.Category) (int64, error) {
	err := u.mysqlDb.Create(category).Error
	return category.ID, err
}

func (u *CategoryRepository) DeleteCategoryByID(id int64) error {
	return u.mysqlDb.Where("id = ?", id).Delete(model.Category{}).Error
}

func (u *CategoryRepository) UpdateCategory(category *model.Category) error {
	return u.mysqlDb.Model(category).Update(&category).Error
}

func (u *CategoryRepository) FindAll() ([]model.Category, error) {
	categories := make([]model.Category, 0)
	return categories, u.mysqlDb.Find(categories).Error
}

func (u *CategoryRepository) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	categories := make([]model.Category, 0)
	err := u.mysqlDb.Where("level = ?", level).Find(categories).Error
	return categories, err
}

func (u *CategoryRepository) FindCategoryByParent(parent int64) ([]model.Category, error) {
	categories := make([]model.Category, 0)
	err := u.mysqlDb.Where("parent = ?", parent).Find(categories).Error
	return categories, err
}

func (u *CategoryRepository) FindCategoryByName(name string) (*model.Category, error) {
	category := &model.Category{}
	err := u.mysqlDb.Where("name = ?", name).Find(category).Error
	return category, err
}
