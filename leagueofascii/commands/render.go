package commands

import (
	"errors"
	"fmt"
	"image"

	"github.com/bloiseleo/leagueofascii/leagueofascii"
	"github.com/bloiseleo/leagueofascii/leagueofascii/helpers"
	"github.com/bloiseleo/leagueofascii/leagueofascii/riot/champions"
)

type RenderCommandOptions struct {
	Champion    string
	Resize      bool
	Width       int
	Height      int
	SquareAsset bool
	Color       bool
}

func getChampion(options RenderCommandOptions) (*champions.Champion, error) {
	if options.Champion == "" {
		return nil, errors.New("champion must be provided")
	}
	champion, err := champions.GetChampion(options.Champion)
	return champion, err
}

func RenderCommand(options RenderCommandOptions) error {
	champion, err := getChampion(options)
	if err != nil {
		return err
	}
	var championImage image.Image
	if options.SquareAsset {
		championImage, err = champions.GetChampionSquareAssets(*champion)
		if err != nil {
			return err
		}
	} else {
		championImage, err = champions.GetChampionLoadingScreen(*champion)
		if err != nil {
			return err
		}
	}
	var art leagueofascii.AscIIArt
	if options.Resize {
		if options.Width <= 0 || options.Height <= 0 {
			return fmt.Errorf("invalid width %v or height %v", options.Width, options.Height)
		}
		championImage = helpers.ResizeImage(championImage, options.Width, options.Height)
	}
	if options.Color {
		art = leagueofascii.CreateAscIIArtWithColors(championImage)
		art.Render()
		return nil
	}
	art = leagueofascii.CreateAscII(championImage)
	art.Render()
	return nil
}
