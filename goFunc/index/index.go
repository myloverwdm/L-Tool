package index

import (
	"LTool/goFunc"
	"encoding/json"
)

type Language struct {
	Name string
}

const languageKey = "Language-xxx"

func GetNowLanguage() string {
	var language Language
	language = goFunc.Get(languageKey, language)
	if language.Name != "" {
		return language.Name
	}
	SetNowLanguage("zh-Hans")
	return "zh-Hans"
}

func SetNowLanguage(language string) {
	marshal, err := json.Marshal(Language{language})
	if err == nil {
		goFunc.Set(languageKey, string(marshal))
		goFunc.GlobalMap[languageKey] = string(marshal)
	}
}
