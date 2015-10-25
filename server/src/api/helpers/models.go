package helpers

import (
	"net/http"
)

// APIItem Model
type APIItem struct {
	ID    int      `json:"-"`
	Cell  []string `json:"cell"`
	Court string   `json:"court"`
}

// CourtItem Model - result that is being finally returned
type CourtItem struct {
	PdfURL  string `json:"pdf_url"`
	Date    string `json:"date"`
	Court   string `json:"court"`
	Title   string `json:"title"`
	Type    string `json:"type"`
	Subject string `json:"subject"`
	Number  string `json:"number"`
}

// APIResponse Model
type APIResponse struct {
	Page    int       `json:"page"`
	Total   int       `json:"total"`
	Records string    `json:"records"`
	Rows    []APIItem `json:"rows"`
}

type HttpResponse struct {
	url       string
	response  *http.Response
	err       error
	courtName string
}
