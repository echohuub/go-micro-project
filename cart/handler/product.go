package handler

import (
	"context"
	"github.com/heqingbao/go-micro-service-cart/common"
	"github.com/heqingbao/go-micro-service-cart/domain/model"
	"github.com/heqingbao/go-micro-service-cart/domain/service"
	"github.com/heqingbao/go-micro-service-cart/proto/cart"
)

type Cart struct {
	CartService service.ICartService
}

func (c *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) error {
	newCart := &model.Cart{}
	err := common.SwapTo(request, newCart)
	if err != nil {
		return err
	}
	id, err := c.CartService.AddCart(newCart)
	if err != nil {
		return err
	}
	response.Id = id
	response.Msg = "添加成功"
	return nil
}

func (c *Cart) CleanCart(ctx context.Context, request *cart.Clean, response *cart.ResponseClean) error {
	err := c.CartService.CleanCart(request.UserId)
	if err != nil {
		return err
	}
	response.Msg = "清空成功"
	return nil
}

func (c *Cart) Incr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	err := c.CartService.IncrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}
	response.Msg = "成功"
	return nil
}

func (c *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	err := c.CartService.DecrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}
	response.Msg = "成功"
	return nil
}

func (c *Cart) DeleteItemById(ctx context.Context, request *cart.CartID, response *cart.Response) error {
	err := c.CartService.DeleteCart(request.Id)
	if err != nil {
		return err
	}
	response.Msg = "成功"
	return nil
}

func (c *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) error {
	cartSlice, err := c.CartService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}

	for _, v := range cartSlice {
		ci := &cart.CartInfo{}
		err := common.SwapTo(v, ci)
		if err != nil {
			break
		}
		response.CartInfo = append(response.CartInfo, ci)
	}
	return nil
}
