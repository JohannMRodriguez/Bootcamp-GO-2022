package service

import (
	"integrador.com/encoder"
	"integrador.com/generator"
)

type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation uint32, err error)
}

type imageGenerator interface {
	BuildAndSaveImage(encodedInformation uint32) error
}

type Information struct {
	UserData       string
	UserHashNumber uint32
}

type Service struct {
	encode   cryptoEncoder
	generate imageGenerator
}

func (s *Service) EncodeInformation(strInformation string) (encodedInformation uint32, err error) {
	encodedInformation, err = encoder.Hash(strInformation)
	return encodedInformation, err
}

func (s *Service) GenerateImage(userHashNumber uint32) error {
	err := generator.GenerateImage(userHashNumber)
	return err
}

func (s *Service) GenerateAndSaveAvatar(information *Information) error {

	userHashNumber, errEncoding := s.EncodeInformation(information.UserData)
	if errEncoding != nil {
		information.UserHashNumber = 0
		return errEncoding
	} else {
		information.UserHashNumber = userHashNumber
		errGeneratingImage := s.GenerateImage(userHashNumber)
		if errGeneratingImage != nil {
			return errGeneratingImage
		} else {
			return nil
		}
	}
}
