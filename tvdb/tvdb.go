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
	apiKey                  string = "5013770E45C5BE20"
	baseURL                 string = "http://thetvdb.com/api"
	getSeriesURL            string = baseURL + "/GetSeries.php?seriesname=%v"
	getSeriesFullDetailsURL string = baseURL + "/" + apiKey + "/series/%v/all/%v.xml"
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

type Episode struct {
	ID                uint32       `xml:"id"`
	EpisodeName       string       `xml:"EpisodeName"`
	Language          string       `xml:"Language"`
	FirstAired        string       `xml:"FirstAired"`
	EpisodeNumber     uint32       `xml:"EpisodeNumber"`
	SeasonNumber      uint32       `xml:"SeasonNumber"`
	Overview          string       `xml:"Overview"`
	CombinedEpisodeNb string       `xml:"Combined_episodenumber"`
	CombinedSeason    string       `xml:"Combined_season"`
	DvdEpNumber       string       `xml:"DVD_episodenumber"`
	DvdSeason         string       `xml:"DVD_season"`
	Director          tvdbPipeList `xml:"Director"`
	GuestStars        tvdbPipeList `xml:"GuestStars"`
	EpImgFlag         string       `xml:"EpImgFlag"`
	ImdbID            string       `xml:"IMDB_ID"`
	ProductionCode    string       `xml:"ProductionCode"`
	Rating            string       `xml:"Rating"`
	RatingCount       string       `xml:"RatingCount"`
	Writer            tvdbPipeList `xml:"Writer"`
	AbsoluteNb        string       `xml:"absolute_number"`
	AirsAfterSeason   string       `xml:"airsafter_season"`
	AirsAfterEpisode  string       `xml:"airsafter_episode"`
	Artwork           string       `xml:"filename"`
	LastUpdated       string       `xml:"lastupdated"`
	SeasonID          uint32       `xml:"seasonid"`
	SeriesID          uint32       `xml:"seriesid"`
	ArtworkAdded      string       `xml:"thumb_added"`
	ArtworkHeight     string       `xml:"thumb_height"`
	ArtworkWidth      string       `xml:"thumb_width"`
}

type tvdbShow struct {
	ID            uint32       `xml:"id"`
	Language      string       `xml:"language"`
	SeriesName    string       `xml:"SeriesName"`
	AliasNames    tvdbPipeList `xml:"AliasNames"`
	Status        string       `xml:"Status"`
	Banner        string       `xml:"banner"`
	Overview      string       `xml:"Overview"`
	Actors        tvdbPipeList `xml:"Actors"`
	AirDate       string       `xml:"Airs_DayOfWeek"`
	AirTime       string       `xml:"Airs_Time"`
	ContentRating string       `xml:"ContentRating"`
	Rating        string       `xml:"Rating"`
	RatingCount   string       `xml:"RatingCount"`
	Runtime       string       `xml:"Runtime"`
	FirstAired    string       `xml:"FirstAired"`
	Genre         tvdbPipeList `xml:"Genre"`
	Added         string       `xml:"added"`
	AddedBy       string       `xml:"addedBy"`
	Fanart        string       `xml:"fanart"`
	LastUpdated   string       `xml:"lastupdated"`
	Poster        string       `xml:"poster"`
	ImdbID        string       `xml:"IMDB_ID"`
	Zap2itID      string       `xml:"zap2it_id"`
	Network       string       `xml:"Network"`
	NetworkID     string       `xml:"NetworkID"`
	Seasons       map[uint32][]*Episode
}

type tvdbShowList struct {
	Series []*tvdbShow `xml:"Series"`
}

type EpisodeList struct {
	Episodes []*Episode `xml:"Episode"`
}

func (show *tvdbShow) GetDetails() (err error) {
	resp, err := http.Get(fmt.Sprintf(getSeriesFullDetailsURL, show.ID, "en"))

	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	seriesList := tvdbShowList{}
	if err = xml.Unmarshal(data, &seriesList); err != nil {
		return
	}
	*show = *seriesList.Series[0]

	episodeList := EpisodeList{}
	if err = xml.Unmarshal(data, &episodeList); err != nil {
		return
	}

	if show.Seasons == nil {
		show.Seasons = make(map[uint32][]*Episode)
	}

	for _, episode := range episodeList.Episodes {
		show.Seasons[episode.SeasonNumber] = append(show.Seasons[episode.SeasonNumber], episode)
	}

	return
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
