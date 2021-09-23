package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"gokitdemo3/userdemo/services"
)

type UserEndpoints struct {
	LoginEndpoint endpoint.Endpoint
	RegisterEndpoint endpoint.Endpoint
}

type LoginRequest struct {

	Email string
	Password string

}

type LoginResponse struct {

	UserId int64	`json:"user_id"`
	UserName string	`json:"user_name"`
	Email string	`json:"email"`
	Err string 	`json:"err"`
}

type RegisterRequest struct {

	UserName string
	Email string
	Password string
}

type RegisterResponse struct {

	UserId int64	`json:"user_id"`
	UserName string `json:"user_name"`
	Email string `json:"email"`
	Err string  `json:"err"`
}

func MakeLoginEndpoint(userService services.UserService)endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:=request.(LoginRequest)
		loginResponse := &LoginResponse{}
		userInfo, err :=userService.Login(ctx,req.Email,req.Password)
		if userInfo==nil || err != nil {
			return loginResponse,err
		}
		return &LoginResponse{
			UserId:   userInfo.UserId,
			UserName: userInfo.UserName,
			Email:    userInfo.Email,
			Err:      userInfo.Err,
		},err
	}
}

func MakeRegisterEndpoint(userService services.UserService)endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req:= request.(RegisterRequest)
		registerResponse := &RegisterResponse{}
		userInfo,err := userService.Register(ctx,req.UserName,req.Email,req.Password)
		if userInfo==nil || err != nil {
			return registerResponse,err
		}
		return &RegisterResponse{
			UserId:   userInfo.UserId,
			UserName: userInfo.UserName,
			Email:    userInfo.Email,
			Err:      userInfo.Err,
		},err
	}
}