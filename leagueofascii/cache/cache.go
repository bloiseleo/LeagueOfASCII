package cache

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"
)

type NodeValue struct {
	V            interface{}
	CreationTime time.Time
}

var mapCache map[string]NodeValue = nil

/*
Get the base path for the cache location
*/
func getBasePath() string {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(cacheDir, "leagueofascii")
}

/*
Get the base path for the file
*/
func getCacheFileLocation() string {
	basePath := getBasePath()
	return filepath.Join(basePath, "cache.json")
}

/*
Creates the file and the directory, if it does not exist, to store the data.
*/
func createCacheFile(fp string) {
	dir, _ := filepath.Split(fp)
	_, err := os.Stat(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	file, err := os.Create(fp)
	if err != nil {
		panic(err)
	}
	file.Close()
}
func loadCacheFromStorage() {
	if mapCache != nil {
		return
	}
	cacheFile := getCacheFileLocation()
	file, err := os.Open(cacheFile)
	if err != nil {
		if os.IsNotExist(err) {
			createCacheFile(cacheFile)
			mapCache = make(map[string]NodeValue)
			return
		}
		panic(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&mapCache)
	if err != nil {
		if errors.Is(err, io.EOF) {
			createCacheFile(cacheFile)
			mapCache = make(map[string]NodeValue)
			return
		}
		panic(err)
	}
}
func createNode(v string) NodeValue {
	return NodeValue{
		V:            v,
		CreationTime: time.Now(),
	}
}

func isCacheValid(v NodeValue) bool {
	t0 := time.Now()
	diff := t0.Sub(v.CreationTime)
	return diff.Hours() < 24
}

/*
Gets a key from the cache and returns the value and if the extraction was successfull.
*/
func GetKeyFromCache(key string) (string, bool) {
	loadCacheFromStorage()
	v, ok := mapCache[key]
	if !ok {
		return "", false
	}
	if !isCacheValid(v) {
		return "", false
	}
	value, ok := v.V.(string)
	if !ok {
		panic("could not convert the value to string")
	}
	return value, true
}

/*
Sets a value into the cache and overwrites the value inside the cache if it exists
*/
func SaveOnCache(k string, v string) {
	loadCacheFromStorage()
	mapCache[k] = createNode(v)
}

/*
Persist the cache to the file system.
*/
func PersistCache() {
	if mapCache == nil {
		return
	}
	file, err := os.OpenFile(getCacheFileLocation(), os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(mapCache)
	if err != nil {
		panic(err)
	}
}
