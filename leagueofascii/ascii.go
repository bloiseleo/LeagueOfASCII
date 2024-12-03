package leagueofascii

import (
	"fmt"
	"image"

	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
	"golang.org/x/image/draw"
)

const asciiChars string = " ░▒▓█"

type AscIIArt interface {
	AscIIMap() [][]string
	Render()
}

type AscIIArtGray struct {
	art [][]string
}

func (art *AscIIArtGray) AscIIMap() [][]string {
	return art.art
}

func (art *AscIIArtGray) Render() {
	for y := range art.art {
		row := art.art[y]
		for x := range row {
			fmt.Printf("%s", row[x])
		}
		fmt.Println()
	}
}

/*
Create an AscII art from the image provided
*/
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
	art := AscIIArtGray{
		art: createMapOfAsciiFromAverage(uintMap),
	}
	return &art
}

/*
Create an AscII art and, before that, it creates another in-memory image that is resized to the newWidth and newHeight
*/
func CreateAscIIAndResize(img image.Image, newWidth, newHeight int) AscIIArt {
	resize := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.CatmullRom.Scale(resize, resize.Rect, img, img.Bounds(), draw.Over, nil)
	return CreateAscII(resize)
}

func createMapOfAsciiFromAverage(brightnessMap [][]uint8) [][]string {
	asciiMap := helpers.CreateStringMap(len(brightnessMap[0]), len(brightnessMap))
	for y := range asciiMap {
		row := asciiMap[y]
		for x := range row {
			asciiMap[y][x] = string(convertBrightnessToAscii(brightnessMap[y][x]))
		}
	}
	return asciiMap
}

func convertBrightnessToAscii(v uint8) rune {
	asciirunes := []rune(asciiChars)
	asciiIndexRelative := int(float64(len(asciirunes)-1) * float64(v) / 255)
	return asciirunes[asciiIndexRelative]
}
