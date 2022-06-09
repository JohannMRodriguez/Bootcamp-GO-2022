package encoder

import (
	"errors"
	"hash/fnv"
)

var errorHashNumber = errors.New("error generating hash number")

// hash function -> gives a hash number for each string passed
func Hash(s string) (uint32, error) {
	h := fnv.New32a()
	h.Write([]byte(s))
	if len([]byte(s)) != 0 {
		return h.Sum32(), nil
	} else {
		return 0, errorHashNumber
	}
}
