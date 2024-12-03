package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bloiseleo/leagueofascii/leagueofascii/cache"
	"github.com/bloiseleo/leagueofascii/leagueofascii/commands"
)

func render(args []string) {
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
		return
	}
	if resize && (newWidth <= 0 || newHeight <= 0) {
		fmt.Println("Error: Resizing must get a new valid width and height")
		render.Usage()
		os.Exit(1)
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
		os.Exit(1)
	}
	os.Exit(0)
}

func main() {
	defer cache.PersistCache()
	flag.Usage = func() {
		fmt.Println("LeagueOfASCII - Welcome to League Of Asc II")
		fmt.Printf("Command expected: %v <command> --flags\n", os.Args[0])
	}
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}
	command := os.Args[1]
	switch command {
	case "render":
		render(os.Args[2:])
	}
}
