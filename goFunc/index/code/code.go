package code

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// 一个代码组, 包括名称和各个代码

type CodeGroup struct {
	Name     string `json:"name"`
	CodeList []Code `json:"codeList"`
}

// 一个代码, 包括类型和代码片段
type Code struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

// 得到所有代码组
const fileName = "data/code.json"

func GetCodeGroup() []CodeGroup {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []CodeGroup{}
	}
	// 将字符串转换为 JSON 对象
	var copyHis []CodeGroup
	err2 := json.Unmarshal(file, &copyHis)
	if err2 != nil {
		// 清空文件
		err := os.Remove(fileName)
		if err != nil {
			return []CodeGroup{}
		}
		return []CodeGroup{}
	}
	return copyHis
}
