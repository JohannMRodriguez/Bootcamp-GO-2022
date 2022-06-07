package main

import (
	"fmt"

	"integrador.com/events/encoder"
	"integrador.com/events/generator"
)

type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation uint32, err error)
}

type imageGenerator interface {
	BuildAndSaveImage(encodedInformation uint32) error
}

type Information struct {
	userInput      string
	userHashNumber uint32
}

func main() {

	var PersonalInfo string

	// get the personal info required
	fmt.Println("Please, enter you email, IP address or public password...")
	fmt.Scan(&PersonalInfo)

	newUser := Information{userInput: PersonalInfo}

	getUserHashNumber, errUserHashNumber := encoder.EncodeInformation(newUser.userInput)
	if errUserHashNumber == nil {
		newUser.userHashNumber = getUserHashNumber
		errUserAvatar := generator.BuildAndSaveImage(newUser.userHashNumber)
		if errUserAvatar != nil {
			fmt.Println("something went wrong creating your avatar :(")
		} else {
			fmt.Println("Your avatar was successfull created!\nWelcome user", newUser.userHashNumber)
		}
	}
}
