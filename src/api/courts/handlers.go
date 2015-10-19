package courts

import (
	"../helpers"
	"github.com/gin-gonic/gin"
)

// IndexAction - Test Action
func IndexAction(context *gin.Context) {
	var results []helpers.Court
	responseCode := 200
	for slug := range helpers.CourtSlugs {
		url := helpers.GenerateAPIUrl(slug)
		if data, code, errors := helpers.MakeAPICall(url); code == 200 && errors == nil {
			results = append(results, data.Rows...)
		} else {
			context.JSON(responseCode, errors)
			break
		}
	}

	context.JSON(responseCode, results)
}

// SearchAction - Recipes list
func SearchAction(context *gin.Context) {
	// url := fmt.Sprintf("%s/search?key=%s", apiURL, apiKEY)
	// data, responseCode := utils.MakeAPICall(url)
	// context.JSON(responseCode, data)
}
