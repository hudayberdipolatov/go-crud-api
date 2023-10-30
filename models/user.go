package models

import (
	"fmt"
	"golang_crud_api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Username string
	FullName string
	Email    string
	Age      string
	gorm.Model
}

func (user User) Migrate() {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&user)
}

func (user User) GetPost(where ...interface{}) User {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return user
	}
	db.First(&user, where...)
	return user
}

func (user User) GetAllPost(where ...interface{}) []User {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var users []User
	db.Find(&users, where...)
	return users
}

func (user User) Create() {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&user)
}

func (user User) Updates(data User) {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Model(&user).Updates(data)
}
func (user User) Delete(where ...interface{}) {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.Delete(&user, where...)
}
