package main

import (
	"fmt"

	"github.com/bloiseleo/leagueofascii/leagueofascii/cache"
	"github.com/bloiseleo/leagueofascii/leagueofascii/riot/champions"
)

func main() {
	defer cache.PersistCache()
	champions := champions.GetAllChampions()
	fmt.Println(champions.Data["Aatrox"])
}
