package main

import (
	"github.com/heqingbao/go-micro-project/category/common"
	"github.com/heqingbao/go-micro-project/category/domain/repository"
	"github.com/heqingbao/go-micro-project/category/handler"
	"github.com/heqingbao/go-micro-project/category/proto/category"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"

	service2 "github.com/heqingbao/go-micro-project/category/domain/service"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		// 设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		// 添加consul作为注册中心
		micro.Registry(consulRegistry),
	)

	// 获取mysql配置，路径中不带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.DB+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// 禁止复表
	db.SingularTable(true)

	srv.Init()

	//rp := repository.NewCategoryRepository(db)
	//rp.InitTable()

	categoryService := service2.NewCategoryService(repository.NewCategoryRepository(db))

	// Register handler
	err = category.RegisterCategoryHandler(srv.Server(), &handler.Category{CategoryService: categoryService})
	if err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
