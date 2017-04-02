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

	cfg := config.Load()
	fmt.Println(cfg.Path[0])
	tvr := tvrenamer.New(cfg.Language, cfg.NameFormat, cfg.NewPath, cfg.Regex, cfg.Move, cfg.Confirm)
	if cfg.SeriesName != "" {
		tvr.SetSeriesName(cfg.SeriesName)
	}
	for _, path := range cfg.Path {
		tvr.Rename(path)
	}
}
