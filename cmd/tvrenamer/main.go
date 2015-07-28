package main

import (
	"fmt"
	"github.com/nomis43/go-tvrenamer"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: tvrenamer [OPTIONS] path")
		os.Exit(1)
	}

	cfg := Config{"en", "{{SeriesName}} - S{{SeasonNb}}E{{EpNumber}} - {{EpName}}", "$HOME/Medialib", "^(.+)[\\.\\ ][Ss]?(\\d{2}|\\d{1})[EeXx]?(\\d{2}).*(\\.\\w{1,4})$", "", tvrenamer.TVDB}

	cfg.Load("$HOME/.config/tvrenamer/tvrenamer.cfg")
	tvr := tvrenamer.TvRenamer{Language: cfg.Language, NameFormatting: cfg.NameFormatting, NewPath: cfg.NewPath, Regex: cfg.Regex}

	log.Println(tvr)
}
