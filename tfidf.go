package tfidf

import (
	"strings"
	"fmt"
	"sync"
	"math"
	"sort"
)

type freqValue float64

func (n freqValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", n)), nil
}

type Frequency struct {
	Word      string  `json:"word"`
	Frequency float64 `json:"frequency,omitempty"`
	TFIDF     float64 `json:"tfidf_score,omitempty"`
}

// Frequencies with interfaces for sorting
type Frequencies []Frequency

func (f Frequencies) Len() int {
	return len(f)
}

func (f Frequencies) Less(i, j int) bool {
	return f[i].TFIDF < f[j].TFIDF
}

func (f Frequencies) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type document struct {
	wordMap     map[string]int
	frequencies Frequencies
}

type Evaluator struct {
	sync.Mutex
	documents []document
	wordsSeen map[string]int
}

func (e *Evaluator) AddDocument(text string) {

	e.Lock()
	if len(e.wordsSeen) == 0 {
		e.wordsSeen = make(map[string]int)
	}
	e.Unlock()

	wordMap := make(map[string]int)
	words := strings.Fields(strings.ToLower(cleanString(text)))

	for i := range words {

		if _, added := wordMap[words[i]]; !added {
			e.Lock()
			e.wordsSeen[words[i]]++
			e.Unlock()
		}

		wordMap[words[i]]++

	}
	docl := document{wordMap:wordMap}

	e.Lock()
	defer e.Unlock()
	e.documents = append(e.documents, docl)
}

func (e *Evaluator) calcTFIDF(num int){
	d := e.documents[num]
	maxFreq := getMaxFreq(d.wordMap)

	frequencies := Frequencies{}
	for i := range d.wordMap {
		// No smoothing
		tf := float64(d.wordMap[i])/float64(maxFreq)
		e.Lock()
		idf := math.Log10(float64(len(e.documents))/ float64(e.wordsSeen[i]))
		e.Unlock()
		frequencies = append(frequencies, Frequency{
			Word:      i,
			Frequency: tf,
			TFIDF: float64(tf) * idf,
		})
	}
	sort.Sort(sort.Reverse(frequencies))
	e.documents[num].frequencies = frequencies
}

func (e *Evaluator) ForDocsCalcTFIDF() {

	for i := range e.documents {
		e.calcTFIDF(i)
	}
}

func (e *Evaluator) GetValues ()[]Frequencies {
	var values []Frequencies
	for i := range e.documents {
		values = append(values, e.documents[i].frequencies)
	}
	return values
}

func getMaxFreq(wordMap map[string]int) int {
	var maxFreq int
	for i := range wordMap {
		if wordMap[i] > maxFreq {
			maxFreq = wordMap[i]
		}
	}
	return maxFreq
}

func cleanString(s string) string {
	charsToClean := []string{".", "!", "?", "\"", ",", "„", "“", "(", ")", "–", ":", "&", "/"}
	for i := range charsToClean {
		s = strings.Replace(s, charsToClean[i], " ", -1)
	}
	return s
}