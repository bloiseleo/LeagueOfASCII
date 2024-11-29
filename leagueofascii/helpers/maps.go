package helpers

func CreateUintMap(width, height uint32) [][]uint8 {
	uintMap := make([][]uint8, height)
	for i := range uintMap {
		uintMap[i] = make([]uint8, width)
	}
	return uintMap
}

func CreateRuneMap(width, height uint32) [][]rune {
	runeMap := make([][]rune, height)
	for i := range runeMap {
		runeMap[i] = make([]rune, width)
	}
	return runeMap
}
