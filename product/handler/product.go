package handler

import (
	"context"
	"github.com/heqingbao/go-micro-project/product/common"
	"github.com/heqingbao/go-micro-project/product/domain/model"
	"github.com/heqingbao/go-micro-project/product/domain/service"
	"github.com/heqingbao/go-micro-project/product/proto/product"
)

type Product struct {
	ProductService service.IProductService
}

func (p *Product) AddProduct(ctx context.Context, request *product.ProductInfo, response *product.ProductResponse) error {
	productInfo := &model.Product{}
	err := common.SwapTo(request, productInfo)
	if err != nil {
		return err
	}
	id, err := p.ProductService.AddProduct(productInfo)
	if err != nil {
		return err
	}
	response.Id = id
	return nil
}

func (p *Product) FindProductByID(ctx context.Context, request *product.RequestID, response *product.ProductInfo) error {
	productInfo, err := p.ProductService.FindProductByID(request.Id)
	if err != nil {
		return err
	}
	err = common.SwapTo(productInfo, response)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) UpdateProduct(ctx context.Context, request *product.ProductInfo, response *product.Response) error {
	productInfo := &model.Product{}
	err := common.SwapTo(request, productInfo)
	if err != nil {
		return err
	}
	err = p.ProductService.UpdateProduct(productInfo)
	if err != nil {
		return err
	}
	response.Msg = "更新成功"
	return nil
}

func (p *Product) DeleteProductByID(ctx context.Context, request *product.RequestID, response *product.Response) error {
	err := p.ProductService.DeleteProduct(request.Id)
	if err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil
}

func (p *Product) FindAllProduct(ctx context.Context, all *product.RequestAll, response *product.AllProduct) error {
	productSlice, err := p.ProductService.FindAllProduct()
	if err != nil {
		return err
	}
	for _, p := range productSlice {
		pi := &product.ProductInfo{}
		err := common.SwapTo(p, pi)
		if err != nil {
			break
		}
		response.ProductInfo = append(response.ProductInfo, pi)
	}
	return nil
}
