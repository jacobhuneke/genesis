package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lloyd/wnram"
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
	cleaned := cleanWord(word)
	e, err := searchForWord(etymologies, cleaned)
	if err != nil {
		return EnglishEtymology{}, fmt.Errorf("unable to locate word %v", word)
	}

	return e, nil
}

// clean word tries to format the word to match entries of the database. Makes singular, present tense, lowercase
func cleanWord(word string) string {
	noSuffix := strings.TrimRight(word, "d")
	return noSuffix
}

func searchWordNet(word string) error {
	wn, err := wnram.New("./dict")
	if err != nil {
		return err
	}
	fmt.Println("successfully loaded wordnet")
	// lookup "yummy"
	if found, err := wn.Lookup(wnram.Criteria{Matching: "create", POS: []wnram.PartOfSpeech{wnram.Verb}}); err != nil {
		return err
	} else {
		// dump details about each matching term to console
		for _, f := range found {
			f.Dump()
			fmt.Println(f.Related(wnram.AlsoSee))
		}
	}
	return nil
}
