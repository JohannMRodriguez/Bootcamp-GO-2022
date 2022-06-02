// package avatar

// import (
// 	"integrador.com/events/encoder"
// 	"integrador.com/events/generator"
// )

// var Users map[uint32]string

// type cryptoEncoder interface {
// 	EncodeInformation(strInformation string) (encodedInformation uint32, err error)
// }

// type imageGenerator interface {
// 	GenerateImage(encodedInformation uint32) error
// }

// // type Service struct {
// // 	encoder   cryptoEncoder
// // 	generator imageGenerator
// // }

// type Information struct {
// 	UserInfo       string
// 	UserHashNumber uint32
// }

// func GenerateHashNumber(information string) error {

// 	getUserHashNumber := []cryptoEncoder{encoder.Information{}}
// 	PersonalHashNumber, errGeneratingHashNumber := getUserHashNumber[0].EncodeInformation(information)
// 	if errGeneratingHashNumber == nil {
// 		return nil
// 	} else {

// 	}
// 	return nil
// }

// func GenerateAndSaveAvatar(information Information) error {
// 	// generate the avatar to the given hash number
// 	getUserIdenticon := []imageGenerator{generator.Information{}}
// 	errGeneratingImage := getUserIdenticon[0].GenerateImage(information.UserHashNumber)
// 	if errGeneratingImage != nil {
// 		return errGeneratingImage
// 	} else {
// 		// errService := service.GenerateAndSaveAvatar(user)
// 		// users[PersonalHashNumber] = generator.UserImage
// 		// fmt.Println(users[0])
// 		return nil
// 	}

// }
