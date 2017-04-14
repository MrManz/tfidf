package tfidf

import ("fmt"
	"strings")

type TFIDF struct {
	Frequencies []Frequency
}

type Frequency struct {
	Word      string  `json:"word"`
	Frequency float64 `json:"frequency"`
}

func cleanString(s string) string{
	charsToClean := []string{".", "!", "?", "\""}
	for i:= range charsToClean {
		s = strings.Replace(s, charsToClean[i], " ", -1)
	}
	return s
}

func NewTFIDF(sentence string) *TFIDF {
	tf := TFIDF{}
	wordMap := make(map[string]int)

	sentence = cleanString(sentence)
	words := strings.Fields(sentence)

	for i:= range words {
		wordMap[words[i]]++
	}

	fmt.Print(wordMap)
	return &tf
}