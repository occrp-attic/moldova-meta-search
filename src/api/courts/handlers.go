package courts

import (
	"../helpers"
	"fmt"
	"github.com/gin-gonic/gin"
)

// IndexAction - Test Action
func IndexAction(context *gin.Context) {
	for key := range helpers.CourtSlugs {
		url := helpers.GenerateAPIUrl(key)
		fmt.Printf("%s ", url)
	}

	context.JSON(200, map[string]interface{}{
		"api_version": []string{"1.0"},
	})
}

// SearchAction - Recipes list
func SearchAction(context *gin.Context) {
	// url := fmt.Sprintf("%s/search?key=%s", apiURL, apiKEY)
	// data, responseCode := utils.MakeAPICall(url)
	// context.JSON(responseCode, data)
}
