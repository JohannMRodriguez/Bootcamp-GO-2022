package main

import (
	"fmt"

	"integrador.com/events/encoder"
	"integrador.com/events/generator"
)

var Users map[uint32]string

func main() {
	// declare my variables
	var PersonalInfo string
	// users := service.getRegisteredUsers()

	// get the personal info required
	fmt.Println("Please, enter you email, IP address or public password...")
	fmt.Scan(&PersonalInfo)

	getUserHashNumber := []encoder.CryptoEncoder{encoder.Information{}}

	// get the personal hash number with encoder package
	PersonalHashNumber, errGeneratingHashNumber := getUserHashNumber[0].EncodeInformation(PersonalInfo)
	if errGeneratingHashNumber != nil {
		fmt.Printf("unexpected error generating hash number %s", errGeneratingHashNumber)
	} else {
		user := generator.Information{UserInfo: PersonalInfo, UserHashNumber: PersonalHashNumber}
		getUserIdenticon := []generator.ImageGenerator{generator.Information{}}
		errGeneratingImage := getUserIdenticon[0].GenerateImage(user.UserHashNumber)
		if errGeneratingImage != nil {
			fmt.Println("something went wrong creating your avatar :(")
		} else {
			fmt.Println("Your avatar was successfull created!\nWelcome user", generator.UserImage)
			// Users[PersonalHashNumber] = generator.UserImage
		}
	}
}
