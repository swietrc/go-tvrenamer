package config

import (
	"code.google.com/p/gcfg"
	"flag"
	"log"
	"os"
)

// Config contains the configuration of the program loaded from the command line and the config file
type Config struct {
	Language       string // Language with which to rename the file(s)
	Move           bool   // Move stores whether to move the files or not
	NameFormatting string // Format
	NewPath        string // Renamed file path
	Regex          string // Regex with which to find the name, the season and the episode number from the file (default can be overriden with -r)
	Path           string // Path of the file to rename
	// REVIEW: switch Config.Path to []string ?
	Scraper uint8
}

// loadFromFile populates Config from a file specified by path
func (c *Config) loadFromFile(path string) {
	filename := os.ExpandEnv(path)
	cfgFile := struct {
		Main struct {
			Language       string
			NameFormatting string
			NewPath        string
			Regex          string
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
	if val := cfgFile.Main.NewPath; val != "" {
		c.NewPath = val
	}
	if val := cfgFile.Main.Regex; val != "" {
		c.Regex = val
	}
}

// loadFromArgs populates Config with the args passed with the command line
func (c *Config) loadFromArgs() {
	flag.StringVar(&c.Language, "l", c.Language, "Language")
	flag.StringVar(&c.NameFormatting, "f", c.NameFormatting, "Format of the new filename")
	flag.StringVar(&c.Regex, "r", c.Regex, "Custom regular expression to match the file information (must match 4 groups)")
	flag.BoolVar(&c.Move, "m", false, "Flag called to move the file to a location specified in the config file")
	flag.Parse()
	c.Path = flag.Arg(0)
}

// Load initializes a Config object from command line args
func Load(path string) *Config {
	// Default values
	c := new(Config)
	c.Language = "en"
	c.NameFormatting = "{{.SeriesName}} - S{{.SeasonNb}}E{{.EpisodeNb}} - {{.EpisodeName}}"
	c.NewPath, _ = os.Getwd()
	c.Regex = "^(.+)[\\.\\ ][Ss]?(\\d{2}|\\d{1})[EeXx]?(\\d{2}).*(\\.\\w{1,4})$"
	c.Scraper = 0
	c.loadFromFile(path)
	c.loadFromArgs()
	return c
}
