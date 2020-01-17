// Generated code. DO NOT EDIT.

package sevensummits

// Peak represents a mountain.
type Peak struct {
	Name string
	// Elevation in meters.
	Elevation int
	Continent string
	// Name of mountain range.
	Range           string
	Country         string
	FirstAscentYear int
}

// Peaks is the Messner version of the Seven Summits.
var Peaks = []Peak{
	{"Mount Everest", 8848, "Asia", "Himalaya", "Nepal/China", 1953},
	{"Aconcagua", 6961, "South America", "Andes", "Argentina", 1897},
	{"Denali", 6194, "North America", "Alaska Range", "United States", 1913},
	{"Kilimanjaro", 5895, "Africa", "", "Tanzania", 1889},
	{"Mount Elbrus", 5642, "Europe", "Caucasus Mountains", "Russia", 1874},
	{"Mount Vinson", 4892, "Antarctica", "Sentinel Range", "", 1966},
	{"Puncak Jaya", 4884, "Australasia", "Sudirman Range", "Indonesia", 1962},
}
var PeaksName = map[string]*Peak{
	"Mount Everest": &Peaks[0],
	"Aconcagua":     &Peaks[1],
	"Denali":        &Peaks[2],
	"Kilimanjaro":   &Peaks[3],
	"Mount Elbrus":  &Peaks[4],
	"Mount Vinson":  &Peaks[5],
	"Puncak Jaya":   &Peaks[6],
}
var PeaksCountry = map[string]*Peak{
	"Nepal/China":   &Peaks[0],
	"Argentina":     &Peaks[1],
	"United States": &Peaks[2],
	"Tanzania":      &Peaks[3],
	"Russia":        &Peaks[4],
	"":              &Peaks[5],
	"Indonesia":     &Peaks[6],
}
var PeaksFirstAscentYear = map[int]*Peak{
	1953: &Peaks[0],
	1897: &Peaks[1],
	1913: &Peaks[2],
	1889: &Peaks[3],
	1874: &Peaks[4],
	1966: &Peaks[5],
	1962: &Peaks[6],
}
