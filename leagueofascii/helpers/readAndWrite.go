package helpers

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

/*
Reads a JPG file and returns an instance of image.Image. This function does not execute any validation in the path being passed. Be sure that the file exists and it's a valid JPEG.
*/
func ReadJpg(path string) (image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	img, err := jpeg.Decode(reader)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func ReadPNG(path string) (image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	img, err := png.Decode(reader)
	if err != nil {
		return nil, err
	}
	return img, err
}

/*
Creates a new empty image using the bounds of the image passed by parameter
*/
func CreateEmptyImageFrom(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	return newImg
}

/*
Writes the new JPEG image to the dst path
*/
func WriteJpg(image image.Image, dst string, options jpeg.Options) error {
	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()
	err = jpeg.Encode(w, image, &options)
	if err != nil {
		return err
	}
	return nil
}

/*
Calculates the width and height of the image
*/
func CalculateWidthAndHeight(img image.Image) (int, int) {
	b := img.Bounds()
	return (b.Max.X - b.Min.X), (b.Max.Y - b.Min.Y)
}

func CreateJpegFromResponse(response *http.Response) (image.Image, error) {
	contentType := response.Header.Get("Content-Type")
	if contentType != "image/jpeg" {
		panic(fmt.Sprintf("no image returned, %v returned", contentType))
	}
	img, err := jpeg.Decode(response.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func CreatePngFromResponse(response *http.Response) (image.Image, error) {
	contentType := response.Header.Get("Content-Type")
	if contentType != "image/png" {
		panic(fmt.Sprintf("no image returned, %v returned", contentType))
	}
	img, err := png.Decode(response.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}
