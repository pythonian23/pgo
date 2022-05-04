package pgo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// API is the API URL
var API = "https://api.politicsandwar.com/graphql"
// Key is the API key
var Key string

func request(query string, out interface{}) error {
	loc := fmt.Sprintf("%s?api_key=%s", API, Key)
	log.Println(loc)
	resp, err := http.PostForm(loc, url.Values{"query": {query}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(contents))
	err = json.Unmarshal(contents, out)
	if err != nil {
		return err
	}
	return nil
}
