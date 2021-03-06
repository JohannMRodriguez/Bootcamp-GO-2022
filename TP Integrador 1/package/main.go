package main

import (
	"fmt"

	"integrador.com/service"
)

func main() {

	var userInfo string

	fmt.Println("Enter your ID, email or password...")
	fmt.Scan(&userInfo)

	newUser := &service.Service{}
	errAvatar := newUser.GenerateAndSaveAvatar(&service.Information{UserData: userInfo})
	if errAvatar != nil {
		fmt.Println("Something went wrong creating your Avatar!")
		fmt.Println(errAvatar)
	} else {
		fmt.Println("Your Avatar was successful created!")
	}
}
