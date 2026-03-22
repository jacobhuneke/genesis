package main

import (
	"fmt"
	"log"

	"github.com/smileart/lemmingo"
)

type config struct {
	etymologies  []EnglishEtymology
	prepositions []string
	lemmingo     *lemmingo.Lemmingo
}

func main() {
	etymologies, err := etymologyJSON()
	if err != nil {
		log.Fatal(err.Error())
	}
	prepositions, err := getTextFromFile("prepositions.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	lem, err := makeLemmingo()
	if err != nil {
		fmt.Println(err.Error())
	}
	c := config{
		etymologies:  etymologies,
		prepositions: prepositions,
		lemmingo:     lem,
	}
	verses1, err := getTextFromFile("genesis1kjv.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	verse1 := getVerse(verses1, 0)

	noPreps := removePrepositions(c.prepositions, verse1)
	_, err = getEtymologiesForVerse(c.etymologies, noPreps)
	if err != nil {
		log.Fatal(err.Error())
	}

	str, _, err := c.lemmingo.Lemma("created", "verb")
	fmt.Println(str)
}
