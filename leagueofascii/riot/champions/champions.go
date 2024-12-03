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

func GetAllChampions() (*ChampionsSummary, error) {
	v, ok := cache.GetKeyFromCache(cache.CHAMPIONS_KEY)
	var response ChampionsSummary
	if ok {
		json.Unmarshal([]byte(v), &response)
		return &response, nil
	}
	resp, err := http.Get(CHAMPIONS_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("no data available at Riot Data Dragon, %v returned", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	cache.SaveOnCache(cache.CHAMPIONS_KEY, string(b))
	return &response, nil
}

func GetChampion(key string) (*Champion, error) {
	var champion ChampionResponse
	url := generateChampionUrl(key)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("no data available at Riot Data Dragon, %v returned", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&champion)
	if err != nil {
		return nil, err
	}
	v, ok := champion.Data[key]
	if !ok {
		return nil, fmt.Errorf("%v was not located inside champion.Data from DataDragon", key)
	}
	return &v, nil
}

func GetChampionLoadingScreen(champion Champion) (image.Image, error) {
	url := generateChampionLoading(champion)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("no data availabe at Riot Data Dragon, %v returned", resp.StatusCode)
	}
	img, err := helpers.CreateJpegFromResponse(resp)
	if err != nil {
		return nil, err
	}
	return img, nil
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
