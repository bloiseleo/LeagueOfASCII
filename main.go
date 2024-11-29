package main

import (
	"fmt"

	"github.com/bloiseleo/leagueofascii/leagueofascii"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

func main() {
	img, err := helpers.ReadJpg("./assets/poro_white.jpg")
	if err != nil {
		panic(err)
	}
	data := leagueofascii.CreateAscII(img)
	for y := range data {
		row := data[y]
		for x := range row {
			fmt.Printf("%c", row[x])
		}
		fmt.Println()
	}
}
