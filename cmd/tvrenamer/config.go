package main

import (
	"code.google.com/p/gcfg"
	"flag"
	"log"
	"os"
)

// Config contains the configuration of the program loaded from the command line and the config file
type Config struct {
	Language       string // Language with which to rename the file(s)
	NameFormatting string // Format
	NewPath        string // Renamed file path
	Regex          string // Regex with which to find the name, the season and the episode number from the file (default can be overriden with -r)
	Path           string // Path of the file to rename
	Scraper        uint8
}

// loadFromFile populates Config from a file specified by path
func (c *Config) loadFromFile(path string) {
	filename := os.ExpandEnv(path)
	cfgFile := struct {
		Main struct {
			Language       string
			NameFormatting string
			NewPath        string
		}
	}{}

	err := gcfg.ReadFileInto(&cfgFile, filename)
	if err != nil {
		log.Println(err)
	}

	if val := cfgFile.Main.Language; val != "" {
		c.Language = val
	}
	if val := cfgFile.Main.NameFormatting; val != "" {
		c.NameFormatting = val
	}
}

func (c *Config) loadFromArgs() {
	argsConf := Config{}
	flag.StringVar(&argsConf.Language, "l", c.Language, "Language")
	flag.StringVar(&argsConf.NameFormatting, "f", c.NameFormatting, "Format of the new filename")
	flag.StringVar(&c.Regex, "r", c.Regex, "Custom regular expression to match the file information (must match 4 groups)")
	flag.Parse()

	c.Path = flag.Arg(0)
}

func (c *Config) Load(path string) {
	c.loadFromFile(path)
	c.loadFromArgs()
}
