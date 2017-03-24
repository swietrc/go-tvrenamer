package tvrenamer

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/swietrc/go-tvrenamer/tvdb"
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
	NewPath        *template.Template
	Regex          *regexp.Regexp
	Scraper        uint8
}

// Rename renames a file based on info parsed from its name.
func (r *TvRenamer) Rename(filepath string) (err error) {
	filepath, err = getAbsolutePath(filepath)
	if err != nil {
		log.Println(err)
	}
	ep, err := r.getEpisode(path.Base(filepath))
	if err != nil {
		return
	}
	var b bytes.Buffer
	var newFilepath string
	pathSNumber, _ := strconv.Atoi(ep.SeasonNb)
	pathEpNumber, _ := strconv.Atoi(ep.EpisodeNb)
	r.NameFormatting.Execute(&b, episode{
		SeriesName:  ep.SeriesName,
		EpisodeNb:   fmt.Sprintf("%02d", pathEpNumber),
		SeasonNb:    fmt.Sprintf("%02d", pathSNumber),
		EpisodeName: ep.EpisodeName,
	})
	newFilename := b.String() + path.Ext(filepath)
	if r.Move {
		b = bytes.Buffer{}

		r.NewPath.Execute(&b, ep)
		log.Println(b.String())
		newFilepath = path.Join(b.String(), newFilename)
		err = os.MkdirAll(path.Dir(newFilepath), 0777)
		if err != nil {
			log.Println("Unable to create dirs to path " + path.Dir(newFilepath))
			log.Fatalln(err)
		}
	} else {
		newFilepath = path.Join(path.Dir(filepath), newFilename)
	}
	log.Println(newFilepath)
	log.Println(newFilename)

	err = os.Rename(filepath, newFilepath)
	if err != nil {
		log.Println(err)
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
func New(language string, nameFormatting string, newPath string, regex string, move bool) *TvRenamer {

	var err error

	tvr := TvRenamer{
		Language: language,
		Move:     move,
	}

	tvr.NameFormatting, err = template.New("Filename formatting").Parse(nameFormatting)
	if err != nil {
		log.Fatal(err)
	}

	tvr.NewPath, err = template.New("Path formatting").Parse(newPath)
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
	if r.CustomName != "" {
		match[1] = r.CustomName
	}
	match[1] = strings.Replace(match[1], ".", " ", -1)
	log.Println(match)
	switch r.Scraper {
	case TVDB:
		seriesList, _ := tvdb.GetSeries(match[1], r.Language)

		if seriesList.Series == nil {
			// return ep, err
			log.Fatalln("ERROR: EPISODE NOT FOUND!")
		}
		fmt.Println("BEGINNING")
		fmt.Println(seriesList.Series[0].SeriesName)
		fmt.Println("END")

		tvdbSeries := seriesList.Series[0]

		tvdbSeries.GetDetails()
		tvdbEp := seriesList.Series[0].Seasons[uint32(seasonNb)][uint32(episodeNb)-1]
		ep = episode{
			SeriesName:  tvdbSeries.SeriesName,
			EpisodeNb:   fmt.Sprint(tvdbEp.EpisodeNumber),
			SeasonNb:    fmt.Sprint(tvdbEp.SeasonNumber),
			EpisodeName: tvdbEp.EpisodeName,
		}
	}
	return ep, nil
}

func getAbsolutePath(p string) (absPath string, err error) {
	absPath = p
	wd, err := os.Getwd()
	if !path.IsAbs(p) {
		absPath = path.Join(wd, p)
	}
	return
}
