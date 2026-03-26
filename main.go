package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/smileart/lemmingo"
)

type config struct {
	etymologies  []EnglishEtymology
	prepositions []string
	lemmingo     *lemmingo.Lemmingo
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err.Error())
	}

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
	ety, err := c.getEtymologiesForVerse(c.etymologies, noPreps)
	if err != nil {
		log.Fatal(err.Error())
	}
	for i, e := range ety {
		fmt.Println(i)
		fmt.Println(e)
	}
}
