package tvdb

import (
	"testing"
)

func TestGetSeries(t *testing.T) {
	seriesList, err := GetSeries("Mr. Robot")

	if err != nil {
		t.Error(err)
	}

	for _, series := range seriesList.Series {
		// t.Log(series)
		if series.SeriesName == "Mr. Robot" {
			return
		}
	}

	t.Error("name received != name sent")
}

func TestGetDetails(t *testing.T) {
	seriesList, err := GetSeries("Mr. Robot")

	if err != nil {
		t.Error(err)
	}

	for _, series := range seriesList.Series {
		err = series.GetFullDetails()
		if err != nil {
			t.Error(err)
		}
	}
}
