package pgo

import (
	"fmt"

	"github.com/pythonian23/pgo/internal/api"
)

// API is the API URL
var API = "https://api.politicsandwar.com/graphql"

// Key is the API key
var Key string

func request(query string) (*api.Data, error) {
	loc := fmt.Sprintf("%s?api_key=%s", API, Key)
	return api.Request(loc, query)
}
