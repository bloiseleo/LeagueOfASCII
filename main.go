package main

import (
	"github.com/bloiseleo/leagueofascii/leagueofascii"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

func main() {
	img, err := helpers.ReadPNG("./assets/Anivia_P.png")
	if err != nil {
		panic(err)
	}
	art := leagueofascii.CreateAscIIAndResize(img, 40, 30)
	art.Render()
}
