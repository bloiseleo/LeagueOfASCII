package leagueofascii

import (
	"fmt"
	"image"

	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

type Pixel struct {
	R uint8
	G uint8
	B uint8
}

type AscIIArtColor struct {
	art [][]string
}

func (art *AscIIArtColor) AscIIMap() [][]string {
	return art.art
}

func (art *AscIIArtColor) Render() {
	for y := range art.art {
		row := art.art[y]
		for x := range row {
			fmt.Print(row[x])
		}
		fmt.Println()
	}
}

func CreateAscIIArtWithColors(img image.Image) AscIIArt {
	b := img.Bounds()
	width, height := helpers.CalculateWidthAndHeight(img)
	pixelMap := helpers.CreateStringMap(width, height)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			pixelMap[y][x] = helpers.Color(uint8(r>>8), uint8(g>>8), uint8(b>>8), "█")
		}
	}
	return &AscIIArtColor{
		art: pixelMap,
	}
}
