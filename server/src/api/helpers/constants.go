package helpers

const (
	// URL - Base courts api url
	URL = "http://instante.justice.md/apps/hotariri_judecata/inst/%COURT_SLUG%/db_hot_grid.php"
	// RowsCount - Count of rows to be queried
	RowsCount = "50"
	// FilterString - Filters JSON
	FilterString = `{"groupOp":"AND","rules":[{"field":"denumire_dosar","op":"cn","data":"%SEARCH_TERM%"}]}`
)

// CourtSlugs
var CourtSlugs = map[string]string{
	"cac":  "Curtea de Apel Chişinău",
	"cab":  "Curtea de Apel Balti",
	"cabe": "Curtea de Apel Bender",
	"cach": "Curtea de Apel Cahul",
	"caco": "Curtea de Apel Comrat",
	"jm":   "Judecătoria Militară",
	"jan":  "Judecătoria Anenii Noi",
	"jba":  "Judecătoria Bălţi",
	"jbs":  "Judecătoria Basarabeasca",
	"jbe":  "Judecătoria Bender",

	"jb":  "Judecătoria Botanica",
	"jbr": "Judecătoria Briceni",
	"jbu": "Judecătoria Buiucani",
	"jch": "Judecătoria Cahul",
	"jc":  "Judecătoria Calarasi",

	"jct": "Judecătoria Cantemir",
	"jca": "Judecătoria Causeni",
	"jcg": "Judecătoria Ceadir-lunga",
	"jcc": "Judecătoria Centru/Grigoriopol",
	"jcm": "Judecătoria Cimislia",
	"jci": "Judecătoria Ciocana",
	"jco": "Judecătoria Comrat",
	"jcr": "Judecătoria Criuleni",
	"jdn": "Judecătoria Donduseni",
	"jdr": "Judecătoria Drochia",
	"je":  "Judecătoria Comerciala de Circumscriptie",
	"jed": "Judecătoria Edinet",
	"jfa": "Judecătoria Falesti",
	"jfl": "Judecătoria Floresti",
	"jgl": "Judecătoria Glodeni",
	"jhn": "Judecătoria Hincesti",
	"jia": "Judecătoria Ialoveni",
	"jlv": "Judecătoria Leova",
	"jns": "Judecătoria Nisporeni",
	"joc": "Judecătoria Ocnita",
	"jor": "Judecătoria Orhei",
	"jrz": "Judecătoria Rezina/Ribnita",
	"jrc": "Judecătoria Riscani Chisinău",
	"jrs": "Judecătoria Riscani",
	"jsi": "Judecătoria Singerei",
	"jsd": "Judecătoria Soldanesti",
	"jsr": "Judecătoria Soroca",
	"jsv": "Judecătoria Stefan_voda",
	"jst": "Judecătoria Straseni",
	"jt":  "Judecătoria Taraclia",
	"jtl": "Judecătoria Telenesti",
	"jun": "Judecătoria Ungheni",
	"jvl": "Judecătoria Vulcanesti",
	"jdb": "Judecătoria Dubasari",
}
