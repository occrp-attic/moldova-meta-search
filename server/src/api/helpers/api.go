package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var r = regexp.MustCompile(`<a href="(.+)?" `)

// GenerateAPIUrl - Generates url for processing court data
func GenerateAPIUrl(slug string) string {
	return strings.Replace(URL, "%COURT_SLUG%", slug, -1)
}

// MakeAPICall Helper function
func MakeAPICall(apiURL string, searchTerm string, courtName string) ([]CourtItem, int, interface{}) {
	var data []CourtItem
	var apiResponse APIResponse
	filterString := strings.Replace(FilterString, "%SEARCH_TERM%", searchTerm, -1)
	v := url.Values{}
	v.Set("rows", fmt.Sprintf("%s", RowsCount))
	v.Add("filters", filterString)
	v.Add("_page", "1")
	v.Add("_search", "true")
	v.Add("sidx", "id")
	v.Add("sord", "asc")
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
		if err := json.Unmarshal(contents, &apiResponse); err != nil {
			fmt.Printf("%s", err.Error())
			return data, 200, nil
			// errors := map[string]interface{}{
			// 	"errors": []string{err.Error()},
			// }
			// // If there was an error return error, that will be returned later
			// return data, 400, errors
		}
		data = processData(apiResponse, courtName, "ca")
		return data, 200, nil
	}
	return data, 400, nil
}

func GetAPIData(searchTerm string) []*HttpResponse {
	ch := make(chan *HttpResponse, len(CourtSlugs))
	responses := []*HttpResponse{}

	for slug, courtName := range CourtSlugs {
		apiURL := GenerateAPIUrl(slug)
		go func(apiURL string) {
			fmt.Printf("Fetching %s \n", apiURL)
			filterString := strings.Replace(FilterString, "%SEARCH_TERM%", searchTerm, -1)
			v := url.Values{}
			v.Set("rows", fmt.Sprintf("%s", RowsCount))
			v.Add("filters", filterString)
			v.Add("_page", "1")
			v.Add("_search", "true")
			v.Add("sidx", "id")
			v.Add("sord", "asc")
			response, err := http.PostForm(apiURL, v)
			ch <- &HttpResponse{apiURL, response, err, courtName, slug}
		}(apiURL)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(CourtSlugs) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

// ParseResponses - parse responses
func ParseResponses(responses []*HttpResponse) []CourtItem {
	var results []CourtItem

	for _, httpResponse := range responses {
		var apiResponse APIResponse
		response := httpResponse.response
		defer response.Body.Close()
		if contents, err := ioutil.ReadAll(response.Body); err == nil {
			if err := json.Unmarshal(contents, &apiResponse); err != nil {
				fmt.Printf("%s", err.Error())
			}
			results = append(results, processData(apiResponse, httpResponse.courtName, httpResponse.slug)...)
		} else {
			fmt.Printf("%s", err.Error())
		}
	}

	return results
}

func processData(data APIResponse, court string, slug string) []CourtItem {
	var results []CourtItem
	for _, value := range data.Rows {
		result := CourtItem{}
		result.Court = court
		result.Date = value.Cell[1]

		// TODO: Extract to a separate function
		pdfURL := value.Cell[0]
		pdfURL = strings.Replace(pdfURL, `\u003c`, `<`, -1)
		pdfURL = strings.Replace(pdfURL, `\u003e`, `>`, -1)
		pdfURL = strings.Replace(pdfURL, `\"`, `"`, -1)
		pdfURL = strings.Replace(pdfURL, `\u0026`, `&`, -1)
		if match := r.FindStringSubmatch(pdfURL); len(match) > 0 {
			result.PdfURL = strings.Replace((PDFBaseURL + match[1]), `%COURT_SLUG%`, slug, -1)
		}

		result.Subject = value.Cell[5]
		result.Title = value.Cell[3]
		result.Type = value.Cell[4]
		result.Number = value.Cell[2]
		results = append(results, result)
	}

	return results
}
