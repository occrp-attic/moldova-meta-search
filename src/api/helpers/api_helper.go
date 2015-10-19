package helpers

import (
	"strings"
)

// GenerateAPIUrl - Generates url for processing court data
func GenerateAPIUrl(slug string) string {
	return strings.Replace(URL, "%COURT_SLUG%", slug, -1)
}
