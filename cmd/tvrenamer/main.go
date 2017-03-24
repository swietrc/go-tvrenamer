package main

import (
	"fmt"
	"os"

	"github.com/swietrc/go-tvrenamer"
	"github.com/swietrc/go-tvrenamer/config"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: tvrenamer [OPTION] FILE")
		os.Exit(1)
	}

	cfg := config.Load("$HOME/.config/tvrenamer/tvrenamer.cfg")

	tvr := tvrenamer.New(cfg.Language, cfg.NameFormatting, cfg.NewPath, cfg.Regex, cfg.Move)

	tvr.Rename(cfg.Path)
}
