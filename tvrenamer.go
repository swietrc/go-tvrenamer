package tvrenamer

const (
	// TVDB represents the tvdb API
	TVDB uint8 = 0
	// TVRAGE represents the tvrage API
	TVRAGE uint8 = 1
)

// TvRenamer stores the configuration used by the rename func
type TvRenamer struct {
	CustomName     string
	Language       string
	NameFormatting string
	NewPath        string
	Regex          string
	Scraper        uint8
}

// Rename a file according to the config object passed as argument
func (r *TvRenamer) Rename(filepath string) (err error) {
	if err != nil {
		return
	}
	return
}
