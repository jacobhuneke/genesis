package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jacobhuneke/genesis/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/smileart/lemmingo"
)

type config struct {
	etymologies  []EnglishEtymology
	prepositions []string
	lemmingo     *lemmingo.Lemmingo
	db           database.Queries
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	dbQueries := database.New(db)

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
		log.Fatal(err.Error())
	}

	c := config{
		etymologies:  etymologies,
		prepositions: prepositions,
		lemmingo:     lem,
		db:           *dbQueries,
	}

	for _, e := range c.etymologies {
		_, er := c.db.GetEtymology(context.Background(), e.Word)
		if er != nil {
			partOfSpeech, _ := getPOS(e.Word)
			params := database.CreateEtymologyParams{
				ID:        uuid.New(),
				Word:      e.Word,
				Etymology: e.Etymology,
				Pos:       partOfSpeech,
			}
			_, er = c.db.CreateEtymology(context.Background(), params)
			if er != nil {
				log.Fatal(er.Error())
			}
		}
	}

	verses1, err := getTextFromFile("genesis1kjv.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	verse1 := getVerse(verses1, 0)

	noPreps := removePrepositions(c.prepositions, verse1)
	_, err = c.getEtymologiesForVerse(c.etymologies, noPreps)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, e := range c.etymologies {
		partOfSpeech, _ := getPOS(e.Word)
		params := database.CreateEtymologyParams{
			ID:        uuid.New(),
			Word:      e.Word,
			Etymology: e.Etymology,
			Pos:       partOfSpeech,
		}
		_, er := c.db.CreateEtymology(context.Background(), params)
		if er != nil {
			log.Fatal(er.Error())
		}
	}
}
