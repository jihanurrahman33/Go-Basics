package user

import (
	"errors"
	"fmt"
	"time"
)
type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}


type Admin struct{
	email string
	password string
	User 
}
func NewAdmin(email,password string)Admin{
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "____",
			createdAt: time.Now(),
		},
	}

}
func (user *User) OutputUserDetails(){
	fmt.Println(user.firstName,user.lastName,user.birthdate)
}

func (user *User) ClearUserName(){
	user.firstName=""
	user.lastName=""


}
func NewUser(firstName,lastName,birthdate string)(*User,error){
	if firstName==""||lastName==""||birthdate==""{
		return nil,errors.New("all fields are required")

	}
	return &User{
		firstName: firstName,
		lastName:lastName,
		birthdate:birthdate,
		createdAt: time.Now(),
	},nil
}