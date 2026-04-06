package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/jdkato/prose/v2"
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
func (c *config) searchForWord(etymologies []EnglishEtymology, word string) (EnglishEtymology, error) {

	for _, e := range etymologies {
		lower := strings.ToLower(e.Word)
		if lower == word {
			return e, nil
		}
	}
	cleaned, err := c.cleanWord(word)
	if err != nil {
		return EnglishEtymology{}, err
	}
	e, err := c.searchForWord(etymologies, cleaned)
	if err != nil {
		return EnglishEtymology{}, fmt.Errorf("unable to locate word %v", word)
	}

	return e, nil
}

// clean word tries to format the word to match entries of the database. Makes singular, present tense, lowercase
func (c *config) cleanWord(word string) (string, error) {
	validLetters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for _, letter := range strings.Split(word, "") {
		if !slices.Contains(validLetters, letter) {
			word = strings.Replace(word, letter, "", 1)
		}
	}
	pos, err := getPOS(word)
	if err != nil {
		return "", err
	}
	newWord, _, err := c.lemmingo.Lemma(word, pos)
	if err != nil {
		return "", err
	}
	return newWord, nil
}

func getPOS(word string) (string, error) {
	d, err := prose.NewDocument(word)
	if err != nil {
		return "", err
	}
	for _, token := range d.Tokens() {
		return token.Tag, nil
	}
	return "", nil
}

/*
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
*/
