package main

import (
	"github.com/heqingbao/go-micro-project/user/domain/repository"
	"github.com/heqingbao/go-micro-project/user/handler"
	"github.com/heqingbao/go-micro-project/user/proto/user"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"

	service2 "github.com/heqingbao/go-micro-project/user/domain/service"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
	)

	srv.Init()

	db, err := gorm.Open("mysql", "root:123456@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	db.SingularTable(true)

	//rp := repository.NewUserRepository(db)
	//rp.InitTable()

	userService := service2.NewUserService(repository.NewUserRepository(db))

	// Register handler
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserService: userService})
	if err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
