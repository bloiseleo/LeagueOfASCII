package leagueofascii

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"math"

	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

const (
	Best_Quality   = 100
	Medium_Quality = 50
	Worst_Quality  = 0
)

func negateColorUint(c uint32) uint8 {
	return uint8(math.MaxUint8 - c>>8)
}

/*
Generates a new Image with all the colors from the image inverted. You can control the quality of the final result using the constants declared:

	Best_Quality   = 100
	Medium_Quality = 50
	Worst_Quality  = 0

But,if you want to provide a custom value, provide an integer value between 0 and 100.
*/
func GenerateNegativeImage(image image.Image, out string, quality int) error {
	if quality < 0 || quality > 100 {
		return errors.New("quality must be between 0 and 100")
	}
	negated := helpers.CreateEmptyImageFrom(image)
	b := negated.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := image.At(x, y)
			r, g, b, a := c.RGBA()
			newRed := negateColorUint(r)
			newBlue := negateColorUint(b)
			newGreen := negateColorUint(g)
			newC := color.RGBA{newRed, newGreen, newBlue, uint8(a)}
			negated.Set(x, y, newC)
		}
	}
	err := helpers.WriteJpg(negated, out, jpeg.Options{
		Quality: quality,
	})
	return err
}
