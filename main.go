package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	etymologies, err := etymologyJSON()
	if err != nil {
		log.Fatal(err.Error())
	}
	verses1, err := getTextFromFile("genesis1esv.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	verse1 := getVerse(verses1, 0)
	prepositions, err := getTextFromFile("prepositions.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	noPreps := removePrepositions(prepositions, verse1)
	word, err := searchForWord(etymologies, noPreps[0])
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(verse1)
	fmt.Println(word.Etymology)
}

// gets the verse specified by the number
func getVerse(verses []string, num int) []string {
	currentVerse := verses[num] //selects it from the slice
	wordsInVerse := strings.Split(currentVerse, " ")
	numWords := len(wordsInVerse)
	wordsInVerse[numWords-1] = strings.Trim(wordsInVerse[numWords-1], ",.'") //removes any trailing punctuation
	return wordsInVerse[1:]                                                  //removes leading verse number and returns a slice of the words in the verse
}

// removes all prepositions as defined by myself in a list (prepositions.txt) from the given verse
func removePrepositions(prepositions, verse []string) []string {
	var result []string
	preps := strings.Join(prepositions, " ")
	for _, word := range verse {
		word = strings.ToLower(word) //lowercase for comparison purposes
		word = strings.TrimRight(word, ",.")
		word = strings.Trim(word, "\"")
		if !strings.Contains(preps, word) {
			result = append(result, word)
		}
	}
	return result
}

func getTextFromFile(url string) ([]string, error) {
	file, err := os.Open(url)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	//gets info to dynamically set file size
	info, err := os.Stat(url)
	if err != nil {
		return []string{}, err
	}
	//stores memory for length of txt file
	size := info.Size()
	versesBytes := make([]byte, size)

	_, err = file.Read(versesBytes) //reads text file
	if err != nil {
		return []string{}, err
	}
	versesStrings := string(versesBytes)             //converts from bytes to strings
	splStrings := strings.Split(versesStrings, "\n") //splits on new line

	return splStrings, nil
}
