package handler

import (
	"context"
	"github.com/heqingbao/go-micro-project/category/common"
	"github.com/heqingbao/go-micro-project/category/domain/model"
	"github.com/heqingbao/go-micro-project/category/domain/service"
	"github.com/heqingbao/go-micro-project/category/proto/category"
	"github.com/prometheus/common/log"
)

type Category struct {
	CategoryService service.ICategoryService
}

func (c *Category) CreateCategory(ctx context.Context, req *category.CategoryRequest, resp *category.CreateCategoryResponse) error {
	newCategory := &model.Category{}
	err := common.SwapTo(req, newCategory)
	if err != nil {
		return err
	}
	id, err := c.CategoryService.AddCategory(newCategory)
	if err != nil {
		return err
	}
	resp.Message = "分类添加成功"
	resp.Id = id
	return nil
}

func (c *Category) UpdateCategory(ctx context.Context, req *category.CategoryRequest, resp *category.UpdateCategoryResponse) error {
	newCategory := &model.Category{}
	err := common.SwapTo(req, newCategory)
	if err != nil {
		return err
	}
	err = c.CategoryService.UpdateCategory(newCategory)
	if err != nil {
		return err
	}
	resp.Message = "分类更新成功"
	return nil
}

func (c *Category) DeleteCategory(ctx context.Context, req *category.DeleteCategoryRequest, resp *category.DeleteCategoryResponse) error {
	err := c.CategoryService.DeleteCategory(req.Id)
	if err != nil {
		return err
	}
	resp.Message = "分类删除成功"
	return nil
}

func (c *Category) FindCategoryByName(ctx context.Context, req *category.FindCategoryByNameRequest, resp *category.CategoryResponse) error {
	newCategory, err := c.CategoryService.FindCategoryByName(req.Name)
	if err != nil {
		return err
	}
	return common.SwapTo(newCategory, resp)
}

func (c *Category) FindCategoryById(ctx context.Context, req *category.FindCategoryByIdRequest, resp *category.CategoryResponse) error {
	newCategory, err := c.CategoryService.FindCategoryByID(req.Id)
	if err != nil {
		return err
	}
	return common.SwapTo(newCategory, resp)
}

func (c *Category) FindCategoryByLevel(ctx context.Context, req *category.FindCategoryByLevelRequest, resp *category.FindAllCategoryResponse) error {
	categorySlice, err := c.CategoryService.FindCategoryByLevel(req.Level)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, resp)
	return nil
}

func (c *Category) FindCategoryByParent(ctx context.Context, req *category.FindCategoryByParentRequest, resp *category.FindAllCategoryResponse) error {
	categorySlice, err := c.CategoryService.FindCategoryByParent(req.Parent)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, resp)
	return nil
}

func (c *Category) FindAllCategory(ctx context.Context, req *category.FindAllCategoryRequest, resp *category.FindAllCategoryResponse) error {
	categorySlice, err := c.CategoryService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, resp)
	return nil
}

func categoryToResponse(categorySlice []model.Category, resp *category.FindAllCategoryResponse) {
	for _, cg := range categorySlice {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		resp.Category = append(resp.Category, cr)
	}
}
