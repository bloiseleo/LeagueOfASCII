package main

import (
	"fmt"

	"github.com/bloiseleo/leagueofascii/leagueofascii/cache"
	"github.com/bloiseleo/leagueofascii/leagueofascii/riot"
)

func main() {
	defer cache.PersistCache()
	latestVersion, _ := riot.GetTheLatestVersionAvailable()
	fmt.Println(latestVersion)
}
