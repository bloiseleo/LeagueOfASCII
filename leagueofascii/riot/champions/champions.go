package champions

import (
	"encoding/json"
	"net/http"
)

const CHAMPIONS_URL = "https://ddragon.leagueoflegends.com/cdn/14.23.1/data/en_US/champion.json"

func GetAllChampions() ChampionsSummary {
	resp, err := http.Get(CHAMPIONS_URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("No 200 returned from champions")
	}
	var response ChampionsSummary
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	return response
}
