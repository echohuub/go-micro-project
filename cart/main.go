package main

import (
	"github.com/heqingbao/go-micro-service-cart/common"
	"github.com/heqingbao/go-micro-service-cart/domain/repository"
	service2 "github.com/heqingbao/go-micro-service-cart/domain/service"
	"github.com/heqingbao/go-micro-service-cart/handler"
	"github.com/heqingbao/go-micro-service-cart/proto/cart"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
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
	// 链路追踪
	tracer, io, err := common.NewTracer("go.micro.service.cart", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer)

	// 数据库设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.DB+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	// 禁止复表
	db.SingularTable(true)

	// 第一次初始化表
	//err = repository.NewCartRepository(db).InitTable()
	//if err != nil {
	//	return
	//}

	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		// 设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8087"),
		// 添加consul作为注册中心
		micro.Registry(consulRegistry),
		// 绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	srv.Init()

	cartService := service2.NewCartService(repository.NewCartRepository(db))

	// Register handler
	err = cart.RegisterCartHandler(srv.Server(), &handler.Cart{CartService: cartService})
	if err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
