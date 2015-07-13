package tvdb

import (
	"fmt"
	"testing"
)

func TestGetSeries(t *testing.T) {
	seriesList, err := GetSeries("Ashes to ashes")

	if err != nil {
		t.Error(err)
	}

	for _, series := range seriesList.Series {
		if series.seriesName == "Ashes to ashes" {
			return
		}
		fmt.Println(series.seriesName)
	}
}
