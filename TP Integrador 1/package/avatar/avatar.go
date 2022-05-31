package avatar

var Test = "teste package avatar"

type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation []byte, err error)
}

type imageGenerator interface {
	ImageGenerator(encodedInformation []byte) error
}

type Service struct {
	encoder   cryptoEncoder
	generator imageGenerator
}

type Information struct {
	userInfo       string
	userHashNumber uint32
}

func (s *Service) GeneratAndSaveAvatar(information Information) error {
	return nil
}
