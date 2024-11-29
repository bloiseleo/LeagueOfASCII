package leagueofascii

import (
	"image"
	"image/color"
	"image/jpeg"

	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

func AvaregeColor(r uint32, g uint32, b uint32) uint8 {
	return uint8(((r >> 8) + (g >> 8) + (b >> 8)) / 3)
}

func GrayScale(image image.Image, out string) error {
	grayscaled := helpers.CreateEmptyImageFrom(image)
	b := image.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pixel := image.At(x, y)
			r, g, b, a := pixel.RGBA()
			gray := AvaregeColor(r, g, b)

			grayscaled.Set(x, y, color.RGBA{
				R: gray,
				G: gray,
				B: gray,
				A: uint8(a),
			})
		}
	}
	return helpers.WriteJpg(grayscaled, out, jpeg.Options{
		Quality: Best_Quality,
	})
}
