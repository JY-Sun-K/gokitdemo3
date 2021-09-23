package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gokitdemo3/userdemo/helper"
	"gokitdemo3/userdemo/pb"
	"google.golang.org/grpc"
	"log"
	"net/http"
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health",func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-type","application/json")
		writer.Write([]byte(`{"status":"ok"}`))
	})
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	//	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	//	w.Header().Set("content-type", "application/json") //返回数据格式是json
	//})

	gwmux := runtime.NewServeMux()

	ctx := context.Background()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredentials())}
	err:=pb.RegisterUserServiceHandlerFromEndpoint(ctx,gwmux,"localhost:8081",opt)
	if err != nil {
		log.Fatalf("从GRPC-GateWay连接GRPC失败, err: %v\n", err)
	}

	mux.Handle("/",gwmux)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}


	err=httpServer.ListenAndServe()
	log.Println(err)
}
