package configuration

type Index struct {
	Name string `json:"name"`

	Route string `json:"route"`

	IconClass string `json:"iconClass"`
}

type Language struct {
	Name string `json:"name"`

	LanguageCode string `json:"languageCode"`
}

var IndexRoute []Index = []Index{
	{"index.util", "/", "fa fa-database"},
}

var LanguageList []Language = []Language{
	{"中文", "zh-Hans"},
	{"English", "en"},
}
