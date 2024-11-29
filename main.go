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
	err = leagueofascii.GenerateNegativeImage(img, "./results/poro_negated.jpg", leagueofascii.Best_Quality)
	if err != nil {
		panic(err)
	}
}
