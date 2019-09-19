package main

import (
	"flag"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

	"log"
	userConfig "mantis-user/config"
	"mantis-user/models"
	userPb "mantis-user/protos"
	"mantis-user/service/impl"
)

func main() {
	configFile := flag.String("f", "./config/config.json", "please use config.json")
	flag.Parse()
	conf := new(userConfig.Config)

	if err := config.LoadFile(*configFile); err != nil {
		log.Fatal(err)
	}

	if err := config.Scan(conf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)

	//etcdRegisty:= etcdv3.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = conf.Etcd.Addrs
	//	// etcdv3.Auth(conf.Etcd.UserName, conf.Etcd.Password)
	//})

	registy := consul.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	engineUser, err := xorm.NewEngine("mysql", "root:@tcp(localhost:3306)/mytest?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	//md := map[string]string{
	//	"vector": "yang",
	//}

	//limit := 1000
	//
	//b := rl.NewBucketWithRate(float64(limit), int64(limit))

	service := micro.NewService(
		micro.Name(conf.Server.Name),
		micro.Registry(registy),
		micro.Version(conf.Version),
		//micro.Metadata(md),
		//micro.Transport(grpc.NewTransport()),
		//micro.WrapHandler(ratelimit.NewHandlerWrapper(b, false)),
	)

	service.Init()
	userModel := models.NewMembersModel(engineUser)
	userRpcServer := impl.NewUserRpcServer(userModel)

	userPb.RegisterUserHandler(service.Server(), userRpcServer)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

