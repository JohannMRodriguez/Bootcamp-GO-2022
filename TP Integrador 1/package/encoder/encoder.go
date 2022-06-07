package encoder

import (
	"hash/fnv"
)

// hash function -> gives a hash number for each string passed
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func EncodeInformation(strInformation string) (encodedInformation uint32, Err error) {
	encodedInformation = hash(strInformation)
	return encodedInformation, nil
}
