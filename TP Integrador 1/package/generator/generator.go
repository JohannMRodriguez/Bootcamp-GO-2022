package generator

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func GenerateImage(encodedInformation uint32) error {

	UserImage := fmt.Sprint(encodedInformation) + ".png"

	width := 50
	height := 50

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	rand.Seed(time.Now().Unix())
	randomColor := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 0xff}
	lightGrey := color.RGBA{236, 236, 232, 0xff}

	// predefine the identicons layouts
	identicon1 := []int{1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1}
	identicon2 := []int{0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0}
	identicon3 := []int{0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0}
	identicon4 := []int{1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1}
	identicon5 := []int{1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0}
	identicon6 := []int{0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0}
	identicon7 := []int{1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1}
	identicon8 := []int{1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1}
	identicon9 := []int{0, 1, 0, 1, 0, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1}
	identicon10 := []int{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0}
	identicon11 := []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0}
	identicon12 := []int{1, 1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0}
	identicons := [12][]int{identicon1, identicon2, identicon3, identicon4, identicon5, identicon6, identicon7, identicon8, identicon9, identicon10, identicon11, identicon12}

	rand.Seed(time.Now().Unix())
	randomIndex := rand.Intn(len(identicons))
	randomIdenticon := identicons[randomIndex]

	// Set color for each pixel.
	actualWidth := width / 5
	actualHeight := height / 5
	lastHeight := 0
	lastWidth := 0

	for index := 0; index < 25; index++ {
		switch {
		case index == 0 || index == 1 || index == 2 || index == 3 || index == 4:
			lastHeight, actualHeight = 0, int(height/5)
		case index == 5 || index == 6 || index == 7 || index == 8 || index == 9:
			lastHeight, actualHeight = int(height/5), int((height*2)/5)
		case index == 10 || index == 11 || index == 12 || index == 13 || index == 14:
			lastHeight, actualHeight = int((height*2)/5), int((height*3)/5)
		case index == 15 || index == 16 || index == 17 || index == 18 || index == 19:
			lastHeight, actualHeight = int((height*3)/5), int((height*4)/5)
		case index == 20 || index == 21 || index == 22 || index == 23 || index == 24:
			lastHeight, actualHeight = int((height*4)/5), width
		}
		switch {
		case index == 0 || index == 5 || index == 10 || index == 15 || index == 20:
			lastWidth, actualWidth = 0, int(width/5)
		case index == 1 || index == 6 || index == 11 || index == 16 || index == 21:
			lastWidth, actualWidth = int(width/5), int((width*2)/5)
		case index == 2 || index == 7 || index == 12 || index == 17 || index == 22:
			lastWidth, actualWidth = int((width*2)/5), int((width*3)/5)
		case index == 3 || index == 8 || index == 13 || index == 18 || index == 23:
			lastWidth, actualWidth = int((width*3)/5), int((width*4)/5)
		case index == 4 || index == 9 || index == 14 || index == 19 || index == 24:
			lastWidth, actualWidth = int((width*4)/5), width
		}
		if randomIdenticon[index] == 1 {
			for y := lastHeight; y < actualHeight; y++ {
				for x := lastWidth; x < actualWidth; x++ {
					img.Set(x, y, randomColor)
				}
			}
		} else {
			for y := lastHeight; y < actualHeight; y++ {
				for x := lastWidth; x < actualWidth; x++ {
					img.Set(x, y, lightGrey)
				}
			}
		}
	}
	// Encode as PNG.
	file, _ := os.Create("generator/img/" + UserImage)
	png.Encode(file, img)

	return nil
}
