package main

import (
	"fmt"
	"time"
)
type User struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}
func (user User) outputUserDetails(){
	fmt.Println(user.firstName,user.lastName,user.birthdate)
}

func (user *User) clearUserName(){
	user.firstName=""
	user.lastName=""

}
func main(){
	userFirstName:=getUserData("Please enter your first name: ")
	userLastName:=getUserData("Please enter your last name: ")
	userBirthdate:=getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	

	
	appUser:=User{
		firstName: userFirstName,
		lastName:userLastName,
		birthdate:userBirthdate,
		createdAt: time.Now(),
		
	}
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

}

func getUserData(promptText string)string{
	fmt.Println(promptText)
	var value string
	fmt.Scan(&value)
	return value
}