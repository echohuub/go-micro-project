package handler

import (
	"context"
	"github.com/heqingbao/go-micro-project/user/domain/model"
	"github.com/heqingbao/go-micro-project/user/domain/service"
	"github.com/heqingbao/go-micro-project/user/proto/user"
)

type User struct {
	UserService service.IUserService
}

func (u *User) Register(ctx context.Context, req *user.UserRegisterRequest, resp *user.UserRegisterResponse) error {
	newUser := &model.User{
		UserName:     req.UserName,
		FirstName:    req.UserFirstName,
		HashPassword: req.Pwd,
	}
	_, err := u.UserService.AddUser(newUser)
	if err != nil {
		return err
	}
	resp.Message = "添加成功"
	return nil
}

func (u *User) Login(ctx context.Context, req *user.UserLoginRequest, resp *user.UserLoginResponse) error {
	ok, err := u.UserService.CheckPwd(req.UserName, req.Pwd)
	if err != nil {
		return err
	}
	resp.IsSuccess = ok
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, req *user.UserInfoRequest, resp *user.UserInfoResponse) error {
	userInfo, err := u.UserService.FindUserByName(req.UserName)
	if err != nil {
		return err
	}
	resp = userForResponse(userInfo)
	return nil
}

func userForResponse(userModel *model.User) *user.UserInfoResponse {
	return &user.UserInfoResponse{
		UserId:        userModel.ID,
		UserName:      userModel.UserName,
		UserFirstName: userModel.FirstName,
	}
}
