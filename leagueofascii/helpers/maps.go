package helpers

func CreateUintMap(width, height uint32) [][]uint8 {
	uintMap := make([][]uint8, height)
	for i := range uintMap {
		uintMap[i] = make([]uint8, width)
	}
	return uintMap
}

func CreateStringMap(width, height int) [][]string {
	runeMap := make([][]string, height)
	for i := range runeMap {
		runeMap[i] = make([]string, width)
	}
	return runeMap
}
