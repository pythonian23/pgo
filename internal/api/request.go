package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Request(accessURL string, query string) (obj *Data, err error) {
	obj = &Data{}
	resp, err := http.PostForm(accessURL, url.Values{"query": {query}})
	if err != nil {
		return
	}
	defer resp.Body.Close()
	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var response = &Response{}
	err = json.Unmarshal(contents, response)
	if err != nil {
		return
	}
	obj = &response.Data
	if len(obj.Errors) != 0 {
		var str []string
		for _, e := range obj.Errors {
			str = append(str, e.Error())
		}
		err = fmt.Errorf("graphql errors: %s", strings.Join(str, "; "))
	}
	return
}
