package main

import (
	"github.com/bloiseleo/leagueofascii/leagueofascii"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

func main() {
	img, err := helpers.ReadJpg("./assets/poro_white.jpg")
	if err != nil {
		panic(err)
	}
	art := leagueofascii.CreateAscII(img)
	art.Render()
}
