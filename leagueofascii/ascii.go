package leagueofascii

import (
	"fmt"
	"image"

	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

const asciiChars string = "░▒▓█"

type AscIIArt struct {
	art [][]rune
}

func (art *AscIIArt) Render() {
	for y := range art.art {
		row := art.art[y]
		for x := range row {
			fmt.Printf("%c", row[x])
		}
		fmt.Println()
	}
}

func CreateAscII(image image.Image) AscIIArt {
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
	art := AscIIArt{
		art: createMapOfAsciiFromAverage(uintMap),
	}
	return art
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
