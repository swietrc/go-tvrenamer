package tvrenamer

import (
	"log"
	"regexp"
	"text/template"
)

const (
	// TVDB represents the tvdb API
	TVDB uint8 = 0
	// TVRAGE represents the tvrage API
	TVRAGE uint8 = 1
)

// TvRenamer stores the configuration used by TvRenamer.Rename to rename files
type TvRenamer struct {
	CustomName     string
	Language       string
	NameFormatting *template.Template
	NewPath        string
	Regex          *regexp.Regexp
	Scraper        uint8
}

// Rename a file according to the config object passed as argument
func (r *TvRenamer) Rename(filepath string) (err error) {
	if err != nil {
		return
	}
	return
}

// SetSeriesName sets TvRenamer.CustomName to the given name
func (r *TvRenamer) SetSeriesName(name string) {
	r.CustomName = name
}

// SetScraper sets TvRenamer.Scraper to the given scraper
func (r *TvRenamer) SetScraper(scraper uint8) {
	r.Scraper = scraper
}

// New allocates a new TvRenamer with the given language, formatting, path and regexp
func New(language string, nameFormatting string, newPath string, regex string) *TvRenamer {

	var err error

	tvr := TvRenamer{
		Language: language,
		NewPath:  newPath,
	}

	tvr.NameFormatting = template.New(nameFormatting)
	tvr.Regex, err = regexp.Compile(regex)

	if err != nil {
		log.Fatal(err)
	}

	return &tvr
}
