package transports

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"gokitdemo3/userdemo/endpoints"
	"gokitdemo3/userdemo/pb"
)

type GRPCServer struct {
	login grpc.Handler
	register grpc.Handler
}

func (g *GRPCServer) Login(ctx context.Context,req *pb.LoginRequest)(*pb.LoginResponse,error) {
	_,rep,err:=g.login.ServeGRPC(ctx,req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginResponse),nil
}

func (g *GRPCServer) Register(ctx context.Context,req *pb.RegisterRequest)(*pb.RegisterResponse,error) {
	_,rep,err:=g.register.ServeGRPC(ctx,req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RegisterResponse),nil
}

func MakeGRPCHandler(userEndpoints endpoints.UserEndpoints) pb.UserServiceServer {
	return &GRPCServer{
		login:    grpc.NewServer(
			userEndpoints.LoginEndpoint,
			decodeLoginRequest,
			encodeLoginResponse,
			),
		register: grpc.NewServer(
			userEndpoints.RegisterEndpoint,
			decodeRegisterRequest,
			encodeRegisterResponse,
			),
	}
}

func decodeLoginRequest(_ context.Context, r interface{}) (interface{}, error) {
	return endpoints.LoginRequest{
		Email:    r.(*pb.LoginRequest).Email,
		Password: r.(*pb.LoginRequest).Password,
	},nil
}
func encodeLoginResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoints.LoginResponse)
	return &pb.LoginResponse{
		UserId:   resp.UserId,
		UserName: resp.UserName,
		Email:    resp.Email,
		Err:      resp.Err,
	},nil
}

func decodeRegisterRequest(_ context.Context, r interface{}) (interface{}, error) {
	return endpoints.RegisterRequest{
		UserName: r.(*pb.RegisterRequest).UserName,
		Email:    r.(*pb.RegisterRequest).Email,
		Password: r.(*pb.RegisterRequest).Password,
	},nil
}
func encodeRegisterResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoints.RegisterResponse)
	return &pb.RegisterResponse{
		UserId:   resp.UserId,
		UserName: resp.UserName,
		Email:    resp.Email,
		Err:      resp.Err,
	},nil
}