package dao

type UserDao interface {
	IsExistUser(email string) (*UserDaoImpl,error)
	CreateUser(userName string,hasedPassword string,email string) (int64,error)
}

type UserDaoImpl struct {
	UserId int64
	UserName string
	Email string
	Password string
	Err string
}

func (u *UserDaoImpl)IsExistUser(email string) (*UserDaoImpl,error) {
	user:=User{}
	err:=DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil,err
	}
	if user.UserId==0 {
		return nil,nil
	}
	return &UserDaoImpl{
		UserId:   user.UserId,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
		Err:      "",
	} ,nil
}

func (u *UserDaoImpl)CreateUser(userName string,hasedPassword string,email string) (int64,error) {
	user :=User{
		UserName:  userName,
		Email:     email,
		Password:  hasedPassword,
	}
	err:=DB.Create(&user).Error
	if err != nil {
		return 0,err
	}
	err=DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return 0,err
	}

	return user.UserId,nil
}

