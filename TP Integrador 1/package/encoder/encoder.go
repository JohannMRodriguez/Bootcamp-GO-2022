package encoder

import "hash/fnv"

type CryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation uint64, err error)
}

type Information struct {
	UserInfo       string
	UserHashNumber uint64
}

// hash function -> gives a hash number for each string passed
func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func (UserInfo Information) EncodeInformation(UserInput string) (EncodedInformation uint64, Err error) {
	EncodedInformation = hash(UserInput)
	return EncodedInformation, nil
}
