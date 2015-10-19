package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// GenerateAPIUrl - Generates url for processing court data
func GenerateAPIUrl(slug string) string {
	return strings.Replace(URL, "%COURT_SLUG%", slug, -1)
}

// MakeAPICall Helper function
func MakeAPICall(apiURL string) (APIResponse, int, interface{}) {
	var data APIResponse
	// Do the request
	v := url.Values{}
	v.Set("rows", fmt.Sprintf("%s", RowsCount))
	response, err := http.PostForm(apiURL, v)
	if err != nil {
		errors := map[string]interface{}{
			"errors": []string{err.Error()},
		}
		// If there was an error return error, that will be returned later
		return data, 404, errors
	}
	defer response.Body.Close()
	if contents, err := ioutil.ReadAll(response.Body); err == nil {
		if err := json.Unmarshal(contents, &data); err != nil {
			errors := map[string]interface{}{
				"errors": []string{err.Error()},
			}
			// If there was an error return error, that will be returned later
			return data, 400, errors
		}
		return data, 200, nil
	}
	return data, 400, nil
}
