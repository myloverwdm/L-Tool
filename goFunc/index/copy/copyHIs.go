package copy

import (
	"LTool/goFunc"
	"encoding/json"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

const fileName = "cache/copyHis.json"

type CopyHis struct {
	Data string `json:"data"`
	// 数据的省略形式
	DataOmit string `json:"dataOmit"`
	// 是否进行了省略
	IsOmit bool   `json:"isOmit"`
	Time   string `json:"time"`
}

func SetCopyHis(copyHis []CopyHis) {
	data, err := json.Marshal(copyHis)
	if err != nil {
		return
	}
	// 写入文件
	os.MkdirAll("cache", os.ModePerm)
	err2 := ioutil.WriteFile(fileName, data, 0644)
	if err2 != nil {
		// fmt.Println(err)
	}
}
func GetCopyHis() []CopyHis {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []CopyHis{}
	}
	// 将字符串转换为 JSON 对象
	var copyHis []CopyHis
	err2 := json.Unmarshal(file, &copyHis)
	if err2 != nil {
		// 清空文件
		err := os.Remove(fileName)
		if err != nil {
			return []CopyHis{}
		}
		return []CopyHis{}
	}
	allSettings := goFunc.GetOrDefaultAllSettings()
	i, err := allSettings.CopySetting.MaxCountOneData.Int64()
	maxCount := int(i)
	for i := 0; i < len(copyHis); i++ {
		// 字符串切分
		if utf8.RuneCountInString(copyHis[i].Data) > maxCount {
			runes := []rune(copyHis[i].Data)
			if len(runes) > maxCount {
				copyHis[i].IsOmit = true
				copyHis[i].DataOmit = string(runes[:maxCount]) + "..."
			}
		} else {
			copyHis[i].IsOmit = false
			copyHis[i].DataOmit = copyHis[i].Data
		}
	}
	return copyHis
}
