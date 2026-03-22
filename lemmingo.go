package main

import (
	"path/filepath"

	"github.com/smileart/lemmingo"
)

func makeLemmingo() (*lemmingo.Lemmingo, error) {
	dictAbsPath, err := filepath.Abs("./dict/en.lmm")
	if err != nil {
		return &lemmingo.Lemmingo{}, err
	}
	lem, err := lemmingo.New(dictAbsPath, "", "freeling", false, false, false)
	if err != nil {
		return &lemmingo.Lemmingo{}, err
	}

	return lem, nil
}

/*
func getPOS(s string) (string, bool) {
	mapPos := tagset.MapPos("penn", "en")
	pos, b := mapPos(s)
	return pos, b
}
*/
