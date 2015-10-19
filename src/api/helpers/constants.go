package helpers

const (
	// URL - Base courts api url
	URL = "http://instante.justice.md/apps/hotariri_judecata/inst/%COURT_SLUG%/db_hot_grid.php"
	// RowsCount - Count of rows to be queried
	RowsCount = 1000
)

// CourtSlugs
var CourtSlugs = map[string]string{
	"cac": "Curtea de Apel Chişinău",
	"cab": "Curtea de Apel Balti",
}
