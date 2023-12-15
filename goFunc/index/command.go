package index

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetCommand() []CommandPojo {
	// 读取内置JSON文件
	data, err := ioutil.ReadFile("data/command.json")
	if err != nil {
		log.Fatal(err)
	}
	var result []CommandPojo
	err2 := json.Unmarshal(data, &result)
	if err2 != nil {
		return []CommandPojo{}
	}
	return result
}

type CommandPojo struct {
	Name string `json:"name"`

	Commands []OneCommand `json:"commands"`
}

type OneCommand struct {
	Desc string `json:"desc"`

	Command string `json:"command"`

	Url string `json:"url"`
}
