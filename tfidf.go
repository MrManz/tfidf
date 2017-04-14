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

func NewTFIDF(sentence string) *TFIDF {
	tf := TFIDF{}
	wordMap := make(map[string]int)

	words := strings.Fields(sentence)

	for i:= range words {
		wordMap[words[i]]++
	}

	fmt.Print(wordMap)
	return tf
}