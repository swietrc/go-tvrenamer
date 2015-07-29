package main

import (
	"fmt"
	"github.com/nomis43/go-tvrenamer"
	"github.com/nomis43/go-tvrenamer/config"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: tvrenamer [OPTIONS] path")
		os.Exit(1)
	}

	cfg := config.Load("$HOME/.config/tvrenamer/tvrenamer.cfg")

	tvr := tvrenamer.TvRenamer{Language: cfg.Language, NameFormatting: cfg.NameFormatting, NewPath: cfg.NewPath, Regex: cfg.Regex}

	log.Println(tvr)
}
