package main

import (
	"github.com/bloiseleo/leagueofascii/leagueofascii"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

func main() {
	img, err := helpers.ReadJpg("./assets/poro.jpg")
	if err != nil {
		panic(err)
	}
	err = leagueofascii.GrayScale(img, "./results/poro_grayscale.jpg")
	if err != nil {
		panic(err)
	}
}
