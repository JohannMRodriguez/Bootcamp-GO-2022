package generator

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

var Test = "Generator test"

type ImageGenerator interface {
	GenerateImage(encodedInformation uint64) error
}

type Information struct {
	UserInfo       string
	UserHashNumber uint64
}

func (UserHashNumber Information) GenerateImage(encodedInformation uint64) error {

	width := 400
	height := 500

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	row1 := int(float32(height) / 0.2)
	row2 := int(float32(height) / 0.4)
	row3 := int(float32(height) / 0.6)
	row4 := int(float32(height) / 0.8)

	column1 := int(float32(width) / 0.25)
	column2 := int(float32(width) / 0.50)
	column3 := int(float32(width) / 0.75)

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 20, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if y < row1 {
				switch {
				case x < column1: // upper left quadrant
					img.Set(x, y, color.RGBA{200, 200, 0, 0xff})
				case x >= column1 && x < column2: // lower right quadrant
					img.Set(x, y, color.RGBA{23, 12, 45, 0xff})
				case x >= column2 && x < column3: // lower right quadrant
					img.Set(x, y, color.White)
				case x >= column3: // lower right quadrant
					img.Set(x, y, color.White)
				default:
					img.Set(x, y, color.Black)
				}
			}
			if y >= row1 && y < row2 {
				switch {
				case x < column1: // upper left quadrant
					img.Set(x, y, cyan)
				case x >= column1 && x < column2: // lower right quadrant
					img.Set(x, y, color.White)
				case x >= column2 && x < column3: // lower right quadrant
					img.Set(x, y, color.White)
				case x >= column3: // lower right quadrant
					img.Set(x, y, color.White)
				default:
					img.Set(x, y, color.Black)
				}
			}
			if y >= row2 && y < row3 {
				switch {
				case x < column1: // upper left quadrant
					img.Set(x, y, cyan)
				case x >= column1 && x < column2: // lower right quadrant
					img.Set(x, y, color.White)
				case x >= column2 && x < column3: // lower right quadrant
					img.Set(x, y, color.White)
				case x >= column3: // lower right quadrant
					img.Set(x, y, color.White)
				default:
					img.Set(x, y, color.Black)
				}
			}
			if y >= row3 && y < row4 {
				switch {
				case x < column1: // upper left quadrant
					img.Set(x, y, cyan)
				case x >= column1 && x < column2: // lower right quadrant
					img.Set(x, y, color.White)
				case x >= column2 && x < column3: // lower right quadrant
					img.Set(x, y, color.White)
				case x >= column3: // lower right quadrant
					img.Set(x, y, color.White)
				default:
					img.Set(x, y, color.Black)
				}
			}
			if y >= row4 {
				switch {
				case x < column1: // upper left quadrant
					img.Set(x, y, cyan)
				case x >= column1 && x < column2: // lower right quadrant
					img.Set(x, y, color.White)
				case x >= column2 && x < column3: // lower right quadrant
					img.Set(x, y, color.White)
				case x >= column3: // lower right quadrant
					img.Set(x, y, color.White)
				default:
					img.Set(x, y, color.Black)
				}
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)

	return nil
}
