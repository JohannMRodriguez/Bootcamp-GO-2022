package generator

var Test = "Generator test"

type imageGenerator interface {
	ImageGenerator(encodedInformation []byte) error
}
