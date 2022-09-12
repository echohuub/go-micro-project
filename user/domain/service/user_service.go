package service

import (
	"errors"
	"github.com/heqingbao/go-micro-project/user/domain/model"
	"github.com/heqingbao/go-micro-project/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) error
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, pwd string) (bool, error)
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{UserRepository: userRepository}
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func (u *UserService) AddUser(user *model.User) (id int64, err error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

func (u *UserService) DeleteUser(id int64) error {
	return u.UserRepository.DeleteUserByID(id)
}

func (u *UserService) UpdateUser(user *model.User, isChangePwd bool) error {
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

func (u *UserService) FindUserByName(name string) (*model.User, error) {
	return u.UserRepository.FindUserByName(name)
}

func (u *UserService) CheckPwd(userName string, pwd string) (bool, error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidatePassword(pwd, user.HashPassword)
}

// GeneratePassword 加密用户密码
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword 验证用户密码
func ValidatePassword(password string, hashed string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return false, errors.New("密码对比错误")
	}
	return true, nil
}
