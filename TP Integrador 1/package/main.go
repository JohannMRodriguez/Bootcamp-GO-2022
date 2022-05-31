package main

import (
	"fmt"

	"integrador.com/events/encoder"
	"integrador.com/events/generator"
)

func main() {
	// declare my variables
	var PersonalInfo string

	// get the personal info required
	fmt.Println("Please, enter you email, IP address or public password...")
	fmt.Scan(&PersonalInfo)

	user := []encoder.CryptoEncoder{encoder.Information{}}
	fmt.Println(user[0])

	// get the personal hash number with encoder package
	PersonalHashNumber, errGeneratingHashNumber := user[0].EncodeInformation(PersonalInfo)
	if errGeneratingHashNumber == nil {
		fmt.Println(PersonalHashNumber, errGeneratingHashNumber)
	} else {
		fmt.Printf("unexpected error generating hash number %s", errGeneratingHashNumber)
	}

	// generate the avatar to the given hash number
	userImage := []generator.ImageGenerator{generator.Information{}}
	errGeneratingImage := userImage[0].GenerateImage(PersonalHashNumber)
	if errGeneratingImage != nil {
		fmt.Printf("unexpected error generating icon %s", errGeneratingImage)
	}
}
