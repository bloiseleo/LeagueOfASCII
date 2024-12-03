package helpers

import "fmt"

func Color(red, green, blue uint8, text string) string {
	return fmt.Sprintf("\033[38;2;%v;%v;%vm%v\033[0m", red, green, blue, text)
}
