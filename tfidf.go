package tfidf

import ("strings"
	"sort")

type Frequency struct {
	Word      string  `json:"word"`
	Frequency float64 `json:"frequency"`
}

// Frequencies with interfaces for sorting
type Frequencies []Frequency

func (f Frequencies) Len() int {
	return len(f)
}

func (f Frequencies) Less(i, j int) bool {
	return f[i].Frequency < f[j].Frequency
}

func (f Frequencies) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}


func cleanString(s string) string{
	charsToClean := []string{".", "!", "?", "\"", ",", "„", "“", "(", ")", "–", ":"}
	for i:= range charsToClean {
		s = strings.Replace(s, charsToClean[i], " ", -1)
	}
	return s
}

func GetFrequenciesFromString(text string) Frequencies {

	wordMap := make(map[string]int)
	words := strings.Fields(cleanString(text))

	for i:= range words {
		wordMap[words[i]]++
	}

	var maxFreq int
	for i := range wordMap {
		if wordMap[i] > maxFreq {
			maxFreq = wordMap[i]
		}
	}

	frequencies := Frequencies{}
	for i := range wordMap {
		frequencies = append(frequencies, Frequency{
			Word:      i,
			Frequency: 0.5 * (1 + float64(wordMap[i])/float64(maxFreq)),
		})
	}

	sort.Sort(sort.Reverse(frequencies))

	return frequencies

}