package avatar

type CryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation uint32, err error)
}

type imageGenerator interface {
	ImageGenerator(encodedInformation []byte) error
}

type Service struct {
	encoder   CryptoEncoder
	generator imageGenerator
}

type Information struct {
	userInfo       string
	userHashNumber uint32
}

func (s *Service) GenerateAndSaveAvatar(information Information) error {
	return nil
}
