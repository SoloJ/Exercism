package scale

import (
	"strings"
)

//Scale is what the test expects
func Scale(tonic, interval string) []string {
	scaleType := map[string][]string{
		"sharps": {"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"},
		"flats":  {"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"},
		"NoSoF":  {""},
	}
	//steps := map[string]int{
	//	"m": 1,
	//	"M": 2,
	//}
	scaleSelect := map[string]string{
		"F":        "flats",
		"Bb":       "flats",
		"Eb":       "flats",
		"Ab":       "flats",
		"Db":       "flats",
		"Gb major": "flats",
		"d":        "flats",
		"g":        "flats",
		"c":        "flats",
		"f":        "flats",
		"bb":       "flats",
		"eb":       "flats",
		"minor":    "flats",
		"C":        "sharps",
		"G":        "sharps",
		"D":        "sharps",
		"A":        "sharps",
		"E":        "sharps",
		"B":        "sharps",
		"F# major": "sharps",
		"e":        "sharps",
		"b":        "sharps",
		"f#":       "sharps",
		"c#":       "sharps",
		"g#":       "sharps",
		"d# minor": "sharps",
		"C major":  "NoSorF",
		"a minor":  "NoSorF",
	}
	var noteType string
	var scaleLength int
	var indexToCut int
	switch {
	case (scaleSelect[tonic] == "flats" && interval == ""):
		d := strings.SplitAfter(strings.Join(scaleType["flats"], ","), tonic)
		d1 := d[1]
		d2 := d[0]
		return strings.SplitAfter(d1+d2, " ")

	case (scaleSelect[tonic] == "sharps"):
		noteType = "sharps"
		scaleLength = 12
		answer := make([]string, scaleLength)
		answer[0] = tonic
		sl1 := make([]string, 0)
		for j, k := range scaleType["sharps"] {
			if string(k) == tonic {
				indexToCut = j
			}
		}
		sl1 = scaleType["sharps"][indexToCut:]
		sl1 = append(sl1, scaleType["sharps"][:indexToCut]...)
		for g

	case (scaleSelect[tonic] == "NoSoF"):
		indexTocut
		noteType = "NoSoF"
		scaleLength = 7
		answer := make([]string, scaleLength)
		answer[0] = tonic
		for j, k := range noteType {
			if k == tonic {
				indexToCut = j
			}
		}

	}

	for i := 1; i < scaleLength; i++ {
		answer[i] = scaleType[noteType][i]
	}
	return answer
}
