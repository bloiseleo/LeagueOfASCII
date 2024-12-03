package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/bloiseleo/leagueofascii/leagueofascii/cache"
	"github.com/bloiseleo/leagueofascii/leagueofascii/commands"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
)

func render(args []string) int {
	var championName string
	var help bool
	var resize bool
	var newWidth int
	var newHeight int
	var squareAsset bool
	render := flag.NewFlagSet("render", flag.ExitOnError)
	render.StringVar(&championName, "champion", "", "Name of the champion to create the ART")
	render.BoolVar(&help, "help", false, "Help about render command")
	render.BoolVar(&resize, "resize", false, "Resize the image before rendering")
	render.IntVar(&newWidth, "width", 0, "New Width")
	render.IntVar(&newHeight, "height", 0, "New Height")
	render.BoolVar(&squareAsset, "square", false, "Gets the Square Asset of the Champion")
	err := render.Parse(args)
	if err != nil {
		panic(err)
	}
	render.Usage = func() {
		fmt.Println("Render, by default, the SplashScreen of the champion")
		fmt.Println("Usage of render:")
		render.PrintDefaults()
	}
	if help || len(args) == 0 {
		render.Usage()
		return 0
	}
	if resize && (newWidth <= 0 || newHeight <= 0) {
		fmt.Println("Error: Resizing must get a new valid width and height")
		render.Usage()
		return 1
	}
	err = commands.RenderCommand(commands.RenderCommandOptions{
		Champion:    championName,
		Resize:      resize,
		Width:       newWidth,
		Height:      newHeight,
		SquareAsset: squareAsset,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		render.Usage()
		return 1
	}
	return 0
}

func main() {
	defer cache.PersistCache()
	flag.Usage = func() {
		fmt.Println("LeagueOfASCII - Welcome to League Of Asc II")
		fmt.Println("Commands: ")
		fmt.Println("- render: renderize a champion insde the terminal in format ASCII")
		fmt.Println()
		fmt.Printf("Usage: %v <command> --flags\n", os.Args[0])
	}
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}
	command := os.Args[1]
	defer helpers.MeasureTime(time.Now())
	switch command {
	case "render":
		render(os.Args[2:])
	default:
		flag.Usage()
	}
}
