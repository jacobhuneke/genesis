package main

import (
	"fmt"
	"io"
	"net/http"
)

func OpenFile(word string) ([]byte, error) {
	url := "https://www.etymonline.com/word/" + word

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	fmt.Println(string(body))
	return body, nil
}
