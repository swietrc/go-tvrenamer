package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"

	"github.com/spf13/viper"
)

// Config contains the configuration of the program loaded from the command line and the config file
type Config struct {
	Language   string   // Language with which to rename the file(s)
	Move       bool     // Move stores whether to move the files or not
	NameFormat string   // Format
	NewPath    string   // Renamed file path
	Regex      string   // Regex with which to find the name, the season and the episode number from the file (default can be overriden with -r)
	Path       []string // Path of the file to rename
	// REVIEW: switch Config.Path to []string ?
	Scraper    uint8
	Confirm    bool
	SeriesName string
}

// loadArgs parses command line args and returns non flag arguments
func loadArgs() []string {
	flagSet := pflag.NewFlagSet("tvrenamer flags", pflag.ExitOnError)
	flagSet.String("name", "", "Custom series name")
	flagSet.StringP("language", "l", viper.GetString("language"), "Language")
	flagSet.StringP("format", "f", viper.GetString("format"), "Format of the new filename")
	flagSet.StringP("regex", "r", viper.GetString("regex"), "Custom regular expression to match the file information (must match 4 groups)")
	flagSet.BoolP("move", "m", false, "Flag called to move the file to a location specified in the config file")
	flagSet.BoolP("confirm", "c", false, "Flag called to move the file to a location specified in the config file")
	flagSet.Parse(os.Args[1:])

	viper.BindPFlags(flagSet)
	return flagSet.Args()
}

// setDefaults sets default configuration values
func setDefaults() {
	viper.SetDefault("format", "{{.SeriesName}} - S{{.SeasonNb}}E{{.EpisodeNb}} - {{.EpisodeName}}")
	viper.SetDefault("language", "en")
	// viper.SetDefault("NewPath", os.Getwd())
	viper.SetDefault("regex", "^(.+)[\\.\\ ][Ss]?(\\d{2}|\\d{1})[EeXx]?(\\d{2}).*(\\.\\w{1,4})$")
	viper.SetDefault("scraper", 0)
	viper.SetDefault("new-path", ".")
}

// Load initializes a Config object from command line args
func Load() *Config {
	// viper.Debug()
	c := new(Config)

	// Set default config filepaths
	viper.SetConfigType("toml")
	viper.SetConfigName("tvrenamer")
	viper.AddConfigPath("/etc/tvrenamer/")
	viper.AddConfigPath("$HOME/.tvrenamer")
	viper.AddConfigPath("$HOME/.config/tvrenamer")
	viper.AddConfigPath(".")
	setDefaults()

	// Parse command line args and read configuration
	c.Path = loadArgs()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Load config into struct
	c.Language = viper.GetString("language")
	c.NameFormat = viper.GetString("format")
	c.NewPath = viper.GetString("new-path")
	c.Regex = viper.GetString("regex")
	c.Move = viper.GetBool("move")
	c.Scraper = 0
	c.Confirm = viper.GetBool("confirm")
	c.SeriesName = viper.GetString("name")

	return c
}
