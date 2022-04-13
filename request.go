package pgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var API string = "https://api.politicsandwar.com/graphql"
var Key string

func request(query string, out interface{}) error {
	url := fmt.Sprintf("%s?api_key=%s&query=%s", API, Key, query)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(contents, out)
	if err != nil {
		return err
	}
	return nil
}
