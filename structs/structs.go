package main

import (
	"fmt"

	"example.com/structs/user"
)

func main(){
	userFirstName:=getUserData("Please enter your first name: ")
	userLastName:=getUserData("Please enter your last name: ")
	userBirthdate:=getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	
	
	appUser,err:=user.NewUser(userFirstName,userLastName,userBirthdate)
	if err!=nil{
		fmt.Println(err)
		return 
	}

	admin:=user.NewAdmin("test@example.com","123456")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()


	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promptText string)string{
	fmt.Println(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}