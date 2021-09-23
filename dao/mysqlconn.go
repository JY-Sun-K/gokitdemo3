package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)
var DB *gorm.DB

type User struct {
	UserId int64 `gorm:"primarykey"`
	UserName string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}



func InitDB() error {
	dsn := "root:sjy1999@tcp(127.0.0.1:3306)/automovie?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err

	}
	err=db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	DB=db
	return nil
}
