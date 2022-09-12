package service

import (
	"github.com/heqingbao/go-micro-project/category/domain/model"
	"github.com/heqingbao/go-micro-project/category/domain/repository"
)

type ICategoryService interface {
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
	FindAllCategory() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
}

func NewCategoryService(categoryRepository repository.ICategoryRepository) ICategoryService {
	return &CategoryService{CategoryRepository: categoryRepository}
}

type CategoryService struct {
	CategoryRepository repository.ICategoryRepository
}

func (u *CategoryService) AddCategory(category *model.Category) (id int64, err error) {
	return u.CategoryRepository.CreateCategory(category)
}

func (u *CategoryService) DeleteCategory(id int64) error {
	return u.CategoryRepository.DeleteCategoryByID(id)
}

func (u *CategoryService) UpdateCategory(category *model.Category) error {
	return u.CategoryRepository.UpdateCategory(category)
}

func (u *CategoryService) FindCategoryByID(id int64) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByID(id)
}

func (u *CategoryService) FindAllCategory() ([]model.Category, error) {
	return u.CategoryRepository.FindAll()
}

func (u *CategoryService) FindCategoryByName(name string) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByName(name)
}

func (u *CategoryService) FindCategoryByParent(parent int64) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByParent(parent)
}

func (u *CategoryService) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByLevel(level)
}
