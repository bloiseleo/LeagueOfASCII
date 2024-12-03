package helpers

import (
	"image"

	"golang.org/x/image/draw"
)

func ResizeImage(img image.Image, newWidth, newHeight int) image.Image {
	resize := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.CatmullRom.Scale(resize, resize.Rect, img, img.Bounds(), draw.Over, nil)
	return resize
}
