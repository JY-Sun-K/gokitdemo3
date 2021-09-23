package services

import (
	"context"
	"gokitdemo3/userdemo/common"
	"gokitdemo3/userdemo/dao"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context,email string,password string)(*common.UserInfo,error)
	Register(ctx context.Context,username string, email string,password string)(*common.UserInfo,error)
}

type UserServiceImpl struct {
	userDao dao.UserDao
}

func MakeUserServiceImpl(userDao dao.UserDao) UserService {
	return&UserServiceImpl{userDao: userDao}
}

func (u *UserServiceImpl) Login(ctx context.Context,email string,password string)(*common.UserInfo,error) {
	user,err:= u.userDao.IsExistUser(email)
	if err != nil || user==nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err!=nil {
		return nil,err
	}

	return &common.UserInfo{
		UserId:   user.UserId,
		UserName: user.UserName,
		Email:    user.Email,
		Err:      "",
	},nil


}

func (u *UserServiceImpl) Register(ctx context.Context,username string, email string,password string)(*common.UserInfo,error) {
	user,_:= u.userDao.IsExistUser(email)

	if user != nil {
		return nil, nil
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil,err
	}
	userId, err := u.userDao.CreateUser(username, string(hasedPassword), email)
	if err != nil {
		return nil,err
	}

	return &common.UserInfo{
		UserId: userId,
		UserName: username,
		Email:    email,
		Err:      "",
	},nil
}