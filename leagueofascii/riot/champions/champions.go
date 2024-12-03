package champions

import (
	"encoding/json"
	"fmt"
	"image"
	"net/http"

	"github.com/bloiseleo/leagueofascii/leagueofascii/cache"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
	"github.com/bloiseleo/leagueofascii/leagueofascii/riot"
)

const CHAMPIONS_URL = "https://ddragon.leagueoflegends.com/cdn/14.23.1/data/en_US/champion.json"
const CHAMPION_URL = "https://ddragon.leagueoflegends.com/cdn/14.23.1/data/en_US/champion/%v.json"
const CHAMPION_LOADING_SCREEN_URL = "https://ddragon.leagueoflegends.com/cdn/img/champion/loading/%v.jpg"
const CHAMPION_SQUARE = "https://ddragon.leagueoflegends.com/cdn/%v/img/champion/%v.png"

func generateChampionUrl(champion string) string {
	return fmt.Sprintf(CHAMPION_URL, champion)
}

func generateChampionLoading(champion Champion) string {
	id := champion.Id
	key := fmt.Sprintf("%v_0", id)
	return fmt.Sprintf(CHAMPION_LOADING_SCREEN_URL, key)
}

func generateChampionSquare(champion Champion, version string) string {
	id := champion.Id
	return fmt.Sprintf(CHAMPION_SQUARE, version, id)
}

func GetAllChampions() ChampionsSummary {
	v, ok := cache.GetKeyFromCache(cache.CHAMPIONS_KEY)
	var response ChampionsSummary
	if ok {
		json.Unmarshal([]byte(v), &response)
		return response
	}
	resp, err := http.Get(CHAMPIONS_URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("No 200 returned from champions")
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	cache.SaveOnCache(cache.CHAMPIONS_KEY, string(b))
	return response
}

func GetChampion(key string) Champion {
	var champion ChampionResponse
	url := generateChampionUrl(key)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("no 200 returned for %v", key))
	}
	err = json.NewDecoder(resp.Body).Decode(&champion)
	if err != nil {
		panic(err)
	}
	v, ok := champion.Data[key]
	if !ok {
		panic("no key in data")
	}
	return v
}

func GetChampionLoadingScreen(champion Champion) image.Image {
	url := generateChampionLoading(champion)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("no 200 returned for %v", url))
	}
	img, err := helpers.CreateJpegFromResponse(resp)
	if err != nil {
		panic(err)
	}
	return img
}

func GetChampionSquareAssets(champion Champion) (image.Image, error) {
	lts, err := riot.GetTheLatestVersionAvailable()
	if err != nil {
		return nil, err
	}
	url := generateChampionSquare(champion, lts)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("no 200 returned while looking for champion square assets, but %v returned", resp.StatusCode)
	}
	img, err := helpers.CreatePngFromResponse(resp)
	if err != nil {
		return nil, err
	}
	return img, nil
}
