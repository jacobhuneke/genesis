package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

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
	/*fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	log.Print("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	*/
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

	gen1, err := getTextFromFile("text/genesis1kjv.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := range len(gen1) {
		verse := getVerse(gen1, i)
		noPreps := removePrepositions(c.prepositions, verse)

		for _, word := range noPreps {
			ety, err := c.db.GetEtymology(context.Background(), word)
			if err != nil {
				cleanedWord, _ := c.cleanWord(word)
				ety, err = c.db.GetEtymology(context.Background(), cleanedWord)
				if err != nil {
					err = notInDB(word)
					if err != nil {
						log.Fatal(err.Error())
					}
				}
			}
			fmt.Println(ety)
		}
	}
}
