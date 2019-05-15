package models

import (
	u "awesomeProject/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

func (user *User) Validate() (map[string] interface{}, bool){

	if user.Name == ""{
		return u.Message(false, "Name can't be empty"), false
	}

	temp := &User{}
	err := GetDB().Table("users").Where("name = ?", user.Name).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return u.Message(false, "Connection error. Please retry"), false
	}

	if temp.Name != ""{
		return u.Message(false, "Name already taken"), false
	}

	fmt.Println("Validated!!!!!!!!!!!!!!!!!!")
	return u.Message(true, "Validated!"), true

}

func (user *User) Create() map[string] interface{}{
	if res, ok := user.Validate(); !ok {
		return res
	}

	GetDB().Create(user)

	//if user.Id <= 0{
	//	return u.Message(false, "Failed to create user, connection error lol")
	//}

	response := u.Message(true, "Account created!")

	response["user"] = user

	return response
}

func GetUser(u uint) *User {

	user := &User{}
	GetDB().Table("users").Where("id = ?", u).First(user)
	if user.Name == ""{
		return nil
	}

	return user
}