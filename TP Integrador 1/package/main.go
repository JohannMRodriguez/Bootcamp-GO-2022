package main

import (
	"fmt"

	"integrador.com/events/encoder"
	// "integrador.com/events/generator"
)

func main() {
	// declare my variables
	var PersonalInfo string

	// get the personal info required
	fmt.Println("Please, enter you email, IP address or public password...")
	fmt.Scan(&PersonalInfo)

	user := []encoder.CryptoEncoder{encoder.Information{}}

	// get the personal hash number with encoder package
	PersonalHashNumber, err := user[0].EncodeInformation(PersonalInfo)
	if err == nil {
		fmt.Println(PersonalHashNumber, err)
	} else {
		fmt.Printf("unexpected error generating hash number %s", err)
	}

	// generate the avatar to the given hash number

}
