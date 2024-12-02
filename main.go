package main

import (
	"github.com/bloiseleo/leagueofascii/leagueofascii"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

func main() {
	img, err := helpers.ReadJpg("./assets/poro_videogame.jpeg")
	if err != nil {
		panic(err)
	}
	art := leagueofascii.CreateAscIIAndResize(img, 100, 100)
	art.Render()
}
