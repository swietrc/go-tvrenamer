package tvrenamer

import (
	"github.com/nomis43/go-tvrenamer/config"
	"io/ioutil"
	"testing"
)

func TestGetEpisode(t *testing.T) {
	c := config.Load("")
	c.Path = "$HOME/Heroes.101.LoL-x264.mp4"
	c.NameFormatting = "{{.SeriesName}}.{{.SeasonNb}}x{{.EpisodeNb}}.{{.EpisodeName}}"
	c.NewPath = "/mnt/media/Multimédia/Séries TV/{{.SeriesName}}/Saison {{.SeasonNb}}/"
	c.Move = true
	tvr := New(c.Language, c.NameFormatting, c.NewPath, c.Regex, c.Move)
	// ep, _ := tvr.getEpisode("Heroes.S01E05.Lol-h264.mp4")
	tvr.Rename(c.Path)
}

func TestSetScraper(t *testing.T) {
	tvr := TvRenamer{}
	tvr.SetScraper(TVDB)
	if tvr.Scraper != TVDB {
		t.Fail()
	}
}

func TestSetSeriesName(t *testing.T) {
	str := "Heroes Reborn"
	tvr := TvRenamer{}
	tvr.SetSeriesName(str)
	if tvr.CustomName != str {
		t.Fail()
	}
}

func TestRename(t *testing.T) {
	dir := ioutil.TempDir("", "tvr")

}
