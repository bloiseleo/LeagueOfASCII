package leagueofascii

import (
	"image"

	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

const asciiChars string = "░▒▓█"

func CreateAscII(image image.Image) [][]rune {
	b := image.Bounds()
	height := b.Max.Y - b.Min.Y
	width := b.Max.X - b.Min.X
	uintMap := helpers.CreateUintMap(uint32(width), uint32(height))
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pixel := image.At(x, y)
			r, g, b, _ := pixel.RGBA()
			grayScaleColor := AvaregeColor(r, g, b)
			uintMap[y][x] = grayScaleColor
		}
	}
	return createMapOfAsciiFromAverage(uintMap)
}

func createMapOfAsciiFromAverage(brightnessMap [][]uint8) [][]rune {
	asciiMap := helpers.CreateRuneMap(uint32(len(brightnessMap[0])), uint32(len(brightnessMap)))
	for y := range asciiMap {
		row := asciiMap[y]
		for x := range row {
			asciiMap[y][x] = convertBrightnessToAscii(brightnessMap[y][x])
		}
	}
	return asciiMap
}

func convertBrightnessToAscii(v uint8) rune {
	asciirunes := []rune(asciiChars)
	asciiIndexRelative := int(float64(len(asciirunes)-1) * float64(v) / 255)
	return asciirunes[asciiIndexRelative]
}
