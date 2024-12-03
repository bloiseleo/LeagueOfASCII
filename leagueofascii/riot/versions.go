package riot

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bloiseleo/leagueofascii/leagueofascii/cache"
)

const VERSION_URL string = "https://ddragon.leagueoflegends.com/api/versions.json"

func GetTheLatestVersionAvailable() (string, error) {
	var patchVersion []string
	v, ok := cache.GetKeyFromCache(cache.VERSION_KEY)
	if ok {
		return v, nil
	}
	resp, err := http.Get(VERSION_URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&patchVersion)
	if err != nil {
		return "", err
	}
	if len(patchVersion) == 0 {
		return "", errors.New("error while decoding body")
	}
	v = patchVersion[0]
	cache.SaveOnCache(cache.VERSION_KEY, v)
	return v, nil
}
