package main

import (
	"fmt"
	"log"
)

func main() {
	etymologies, err := etymologyJSON()
	if err != nil {
		log.Fatal(err.Error())
	}

	verses1, err := getTextFromFile("genesis1kjv.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	verse1 := getVerse(verses1, 0)
	prepositions, err := getTextFromFile("prepositions.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	noPreps := removePrepositions(prepositions, verse1)
	fmt.Println(noPreps)
	ety, err := getEtymologiesForVerse(etymologies, noPreps)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(ety)

}
