package main

import (
	pb "mantis-user/protos"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

func main() {


	// 创建服务并添加响应的配置
	service := micro.NewService(
		micro.Name("greeter.client"),
	)

	service.Init()

	// 创建客户端
	greeter := pb.NewUserService("greeter", service.Client())

	rsp,err := greeter.FindByID(context.TODO(),&pb.FindByIDRequest{Id:1})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp)
}
