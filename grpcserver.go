package main

import (
	"fmt"
	"gokitdemo3/userdemo/dao"
	"gokitdemo3/userdemo/endpoints"
	"gokitdemo3/userdemo/helper"
	"gokitdemo3/userdemo/pb"
	"gokitdemo3/userdemo/services"
	"gokitdemo3/userdemo/transports"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	errChan := make(chan error)
	err:=dao.InitDB()
	if err != nil {
		log.Println(err)
		errChan <- err
	}

	userDao :=&dao.UserDaoImpl{}
	userService := services.MakeUserServiceImpl(userDao)
	userEndpoints :=endpoints.UserEndpoints{
		LoginEndpoint:    endpoints.MakeLoginEndpoint(userService),
		RegisterEndpoint: endpoints.MakeRegisterEndpoint(userService),
	}


	r := transports.MakeGRPCHandler( userEndpoints)

	rpcServer:=grpc.NewServer(grpc.Creds(helper.GetServerCredentials()))
	pb.RegisterUserServiceServer(rpcServer,r)

	lis,_:=net.Listen("tcp",":8081")


	go func() {
		errChan <- rpcServer.Serve(lis)
	}()

	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	log.Println(error)
}
