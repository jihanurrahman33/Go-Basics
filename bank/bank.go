package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile="balance.txt"


func main(){
	var accountBalance,err = fileops.GetFloatFromFile(accountBalanceFile,0)
	if err !=nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
	}
	fmt.Println("Welcome to GO Bank!")
	fmt.Println("Reach us 24/7",randomdata.PhoneNumber())
	for {
		presentOptions()
		var choice int 
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
	
	switch choice{
	case 1:
		fmt.Println("Your Balance is ",accountBalance)
	case 2:
		fmt.Print("Your Deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		if depositAmount<=0{
			fmt.Println("Invalid Amount.Must be greater than 0.")
			continue
		}
		accountBalance += depositAmount
		fmt.Println("Balace Updated! New Amount:",accountBalance)
		fileops.WriteFloatToFile(accountBalance,accountBalanceFile)
	case 3:
		fmt.Print("Withdraw Money:")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if withdrawAmount>accountBalance {
			fmt.Println("Withdrawal Amount is greater than Account balance:",accountBalance)
			continue
		}
		if withdrawAmount<=0{
			fmt.Println("Invalid Amount.Must be greater than 0.")
			continue
		}
		accountBalance-=withdrawAmount
		fmt.Println("Withdrawal Succesfull! New Amount:",accountBalance)
		fileops.WriteFloatToFile(accountBalance,accountBalanceFile)
	default:
		fmt.Println("GoodBye!")
		fmt.Println("Thanks for choosing our bank")
		return	
	}
		}
}
