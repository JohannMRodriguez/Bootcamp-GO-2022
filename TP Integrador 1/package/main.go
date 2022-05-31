package main

import (
	"fmt"
	"hash/fnv"

	"integrador.com/events/avatar"
)

// hash function -> gives a hash number for each string passed
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {

	// declare my variables
	var PersonalInfo string
	// userAvatar := &avatar.Information{}

	// get the personal info required
	fmt.Println("Please, enter you email, IP address or public password...")
	fmt.Scan(&PersonalInfo)

	// get the personal hash number with the
	PersonalHashNumber := hash(PersonalInfo)
	fmt.Println(PersonalInfo, PersonalHashNumber, avatar.Test)

	// generate the avatar to the given hash number

}
