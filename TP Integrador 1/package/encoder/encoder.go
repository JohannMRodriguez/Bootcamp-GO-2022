package encoder

import (
	"fmt"
	"hash/fnv"
)

var EncodedInformation uint32

type CryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation uint32, err error)
}

type Information struct {
	userInfo       string
	userHashNumber uint32
}

// hash function -> gives a hash number for each string passed
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	fmt.Println(s)
	return h.Sum32()
}

func (info Information) EncodeInformation(UserInput string) (EncodedInformation uint32, Err error) {
	EncodedInformation = hash(UserInput)
	return EncodedInformation, nil
}
