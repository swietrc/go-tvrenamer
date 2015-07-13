package tvdb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiKey             string = "5013770E45C5BE20"
	baseURL            string = "http://thetvdb.com/api"
	getSeriesURL       string = baseURL + "/GetSeries.php?seriesname=%v"
	getSeriesByIDURL   string = baseURL + "/api/" + apiKey + "/series/%v/%v.xml"
	getSeriesDetailURL string = baseURL + "/" + apiKey + "/series/%v/%v.xml"
)

type tvdbPipeList []string

func (pipe *tvdbPipeList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	content := ""

	if err = d.DecodeElement(&content, &start); err != nil {
		return
	}

	*pipe = strings.Split(content, "|")

	return
}

type tvdbShowList struct {
	Series []tvdbShow `xml:"Series"`
}

type tvdbShow struct {
	ID            uint32       `xml:"id"`
	language      string       `xml:"language"`
	seriesName    string       `xml:"SeriesName"`
	aliasNames    tvdbPipeList `xml:"AliasNames"`
	status        string       `xml:"Status"`
	banner        string       `xml:"banner"`
	overview      string       `xml:"overview"`
	airDate       string       `xml:"Airs_DayOfWeek"`
	airTime       string       `xml:"Airs_Time"`
	contentRating string       `xml:"ContentRating"`
	rating        string       `xml:"Rating"`
	ratingCount   uint32       `xml:"RatingCount"`
	runtime       uint32       `xml:"Runtime"`
	firstAired    string       `xml:"FirstAired"`
	genre         tvdbPipeList `xml:"Genre"`
	added         string       `xml:"added"`
	addedBy       string       `xml:"addedBy"`
	fanart        string       `xml:"fanart"`
	lastUpdated   string       `xml:"lastupdated"`
	posters       string       `xml:"posters"`
	imdbID        string       `xml:"IMDB_ID"`
	zap2itID      string       `xml:"zap2it_id"`
	network       string       `xml:"Network"`
	networkID     uint32       `xml:"NetworkID"`
}

type tvdbEpisode struct {
	ID                uint32       `xml:"id"`
	episodeName       string       `xml:"EpisodeName"`
	language          string       `xml:"Language"`
	firstAired        string       `xml:"FirstAired"`
	episodeNumber     uint32       `xml:"EpisodeNumber"`
	seasonNumber      uint32       `xml:"SeasonNumber"`
	overview          string       `xml:"Overview"`
	combinedEpisodeNb uint32       `xml:"Combined_episodenumber"`
	combinedSeason    uint32       `xml:"Combined_season"`
	dvdEpNumber       uint32       `xml:"DVD_episodenumber"`
	dvdSeason         uint32       `xml:"DVD_season"`
	director          tvdbPipeList `xml:"Director"`
	guestStars        tvdbPipeList `xml:"GuestStars"`
	epImgFlag         uint32       `xml:"EpImgFlag"`
	imdbID            string       `xml:"IMDB_ID"`
	productionCode    string       `xml:"ProductionCode"`
	rating            string       `xml:"Rating"`
	ratingCount       uint32       `xml:"RatingCount"`
	writer            tvdbPipeList `xml:"Writer"`
	absoluteNb        uint32       `xml:"absolute_number"`
	airsAfterSeason   uint32       `xml:"airsafter_season"`
	airsAfterEpisode  uint32       `xml:"airsafter_episode"`
	artwork           string       `xml:"filename"`
	lastUpdated       string       `xml:"lastupdated"`
	seasonID          uint32       `xml:"seasonid"`
	seriesID          uint32       `xml:"seriesid"`
	artworkAdded      string       `xml:"thumb_added"`
	artworkHeight     uint32       `xml:"thumb_height"`
	artworkWidth      uint32       `xml:"thumb_width"`
}

func GetSeries(name string) (seriesList tvdbShowList, err error) {
	resp, err := http.Get(fmt.Sprintf(getSeriesURL, url.QueryEscape(name)))

	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	err = xml.Unmarshal(data, &seriesList)

	return
}
