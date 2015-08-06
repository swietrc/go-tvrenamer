package tvrenamer

import (
	"bytes"
	"fmt"
	"github.com/nomis43/go-tvrenamer/tvdb"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"text/template"
)

const (
	// TVDB represents the tvdb API
	TVDB uint8 = 0
	// TVRAGE represents the tvrage API
	// TVRAGE uint8 = 1
)

type episode struct {
	SeriesName  string
	EpisodeNb   string
	SeasonNb    string
	EpisodeName string
}

// TvRenamer stores the configuration used by TvRenamer.Rename to rename files
type TvRenamer struct {
	CustomName     string
	Language       string
	Move           bool
	NameFormatting *template.Template
	NewPath        string
	Regex          *regexp.Regexp
	Scraper        uint8
}

// Rename a file according to the config object passed as argument
func (r *TvRenamer) Rename(filepath string) (err error) {
	episode, err := r.getEpisode(filepath)
	if err != nil {
		return
	}
	var b bytes.Buffer
	var newFilepath string
	r.NameFormatting.Execute(&b, episode)
	newFilename := b.String() + path.Ext(filepath)
	if r.Move {
		newFilepath = path.Join(path.Base(filepath), newFilename)
	} else {
		newFilepath = path.Join(r.NewPath, newFilename)
	}
	os.Rename(filepath, newFilepath)
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
func New(language string, nameFormatting string, newPath string, regex string, move bool) *TvRenamer {

	var err error

	tvr := TvRenamer{
		Language: language,
		NewPath:  newPath,
		Move:     move,
	}

	tvr.NameFormatting, err = template.New("Filename formatting").Parse(nameFormatting)
	if err != nil {
		log.Fatal(err)
	}

	tvr.Regex, err = regexp.Compile(regex)
	if err != nil {
		log.Fatal(err)
	}

	return &tvr
}

func (r *TvRenamer) getEpisode(filepath string) (episode, error) {
	ep := episode{}
	match := r.Regex.FindStringSubmatch(path.Base(filepath))
	episodeNb, _ := strconv.ParseUint(match[3], 10, 8)
	seasonNb, _ := strconv.ParseUint(match[2], 10, 8)
	switch r.Scraper {

	case TVDB:
		if r.CustomName != "" {
			match[1] = r.CustomName
		}
		seriesList, err := tvdb.GetSeries(match[1], r.Language)
		if len(seriesList.Series) == 0 {
			return ep, err
		}
		tvdbSeries := seriesList.Series[0]
		tvdbSeries.GetDetails()
		tvdbEp := seriesList.Series[0].Seasons[uint32(seasonNb)][uint32(episodeNb)-1]
		ep = episode{
			SeriesName:  tvdbSeries.SeriesName,
			EpisodeNb:   fmt.Sprintf("%02d", tvdbEp.EpisodeNumber),
			SeasonNb:    fmt.Sprintf("%02d", tvdbEp.SeasonNumber),
			EpisodeName: tvdbEp.EpisodeName,
		}
	}
	return ep, nil
}
