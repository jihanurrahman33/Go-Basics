package main

import "fmt"
func main(){


	var items map[string]string =map[string]string{
		"item_1":"Apple",
		"item_2":"Orange",
	}
	fmt.Println(items["item_1"])
}