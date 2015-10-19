package helpers

// Court Model
type Court struct {
	ID   int      `json:"-"`
	Cell []string `json:"cell"`
}

// APIResponse Model
type APIResponse struct {
	Page    int     `json:"page"`
	Total   int     `json:"total"`
	Records string  `json:"records"`
	Rows    []Court `json:"rows"`
}
