package main

import (
	"fmt"
	"github.com/nomis43/go-tvrenamer"
	"github.com/nomis43/go-tvrenamer/config"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: tvrenamer [OPTIONS] path")
		os.Exit(1)
	}

	cfg := config.Load("$HOME/.config/tvrenamer/tvrenamer.cfg")

	tvr := tvrenamer.New(cfg.Language, cfg.NameFormatting, cfg.NewPath, cfg.Regex, cfg.Move)

	tvr.Rename(cfg.Path)
}
