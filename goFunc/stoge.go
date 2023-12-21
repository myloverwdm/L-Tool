package goFunc

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	GlobalMap map[string]string
)

const cacheFileName = "cache/cacheLS.json"

type cache struct{}

type LTResponse struct {
	Success  bool   `json:"success"`
	Data     string `json:"data"`
	ErrorMsg string `json:"errorMsg"`
}

func Get[T any](key string, classS T) T {
	cache := GlobalMap[key]
	if cache != "" {
		err := json.Unmarshal([]byte(cache), &classS)
		if err != nil {
			// fmt.Println("缓存数据解析失败", err)
			return classS
		}
		return classS
	}
	allCache := ReadAllCache()
	err := json.Unmarshal([]byte(allCache[key]), &classS)
	if err != nil {
		return classS
	}
	// 输出结构体的内容
	GlobalMap[key] = allCache[key]
	return classS
}

func Set(key string, value string) {
	allCache := ReadAllCache()
	allCache[key] = value
	marshal, err := json.Marshal(allCache)
	if err == nil {
		Write(string(marshal))
		GlobalMap[key] = value
	}
}

func ReadAllCache() map[string]string {
	file, err := ioutil.ReadFile(cacheFileName)
	if err != nil {
		return map[string]string{}
	}
	// 将字符串转换为 JSON 对象
	var jsonObj map[string]string
	err2 := json.Unmarshal(file, &jsonObj)
	if err2 != nil {
		// 清空文件
		Remove()
		return map[string]string{}
	}
	return jsonObj
}

func Remove() {
	err := os.Remove(cacheFileName)
	if err != nil {
		return
	}
}

func Write(data string) {
	// fmt.Println("写入缓存")
	content := []byte(data)
	// 写入文件
	err := ioutil.WriteFile(cacheFileName, content, 0644)
	if err != nil {
		// fmt.Println(err)
	}
}
