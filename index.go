package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type EnglishEtymology struct {
	Word            string   `json:"word"`
	Pos             string   `json:"pos"`
	Crossreferences []string `json:"crossreferences"`
	Etymology       string   `json:"etymology"`
	Years           []int    `json:"years"`
}

// reads all the words from etymonline's index to json structs and saves in main
func etymologyJSON() ([]EnglishEtymology, error) {
	var words []EnglishEtymology
	data, err := os.ReadFile("index.json")
	if err != nil {
		return []EnglishEtymology{}, err
	}
	err = json.Unmarshal(data, &words)
	if err != nil {
		return []EnglishEtymology{}, err
	}
	return words, nil
}

// looks for given word in etymology dictionary
func searchForWord(etymologies []EnglishEtymology, word string) (EnglishEtymology, error) {
	for _, e := range etymologies {
		lower := strings.ToLower(e.Word)
		if lower == word {
			return e, nil
		}
	}
	return EnglishEtymology{}, fmt.Errorf("unable to locate word")
}
