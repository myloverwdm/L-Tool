package goFunc

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	GlobalMap map[string]string

	SettingMapCache SystemSetting
)

const cacheFileName = "cache/cache.json"
const settingFileName = "cache/setting.json"

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
		Write(string(marshal), cacheFileName)
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

func Write(data string, fileName string) {
	// fmt.Println("写入缓存")
	content := []byte(data)
	// 写入文件
	os.MkdirAll("cache", os.ModePerm)
	err := ioutil.WriteFile(fileName, content, 0644)
	if err != nil {

	}
}

func GetOrDefaultAllSettings() SystemSetting {
	if SettingMapCache != (SystemSetting{}) {
		return SettingMapCache
	}
	// 读取文件
	file, err := ioutil.ReadFile(settingFileName)
	if err != nil {
		return BuildDefaultSetting()
	}
	// file转化为SystemSetting
	err2 := json.Unmarshal(file, &SettingMapCache)
	if err2 == nil {
		return SettingMapCache
	}
	return BuildDefaultSetting()
}

func BuildDefaultSetting() SystemSetting {
	SettingMapCache = SystemSetting{
		Language: "zh-Hans",
		CopySetting: CopySetting{
			MaxCount:        "100",
			MaxCountOneData: "100",
		},
	}
	return SettingMapCache
}

func UpdateSystemSettings(setting string) {
	var sysSet SystemSetting
	err2 := json.Unmarshal([]byte(setting), &sysSet)
	if err2 != nil {

	}
	marshal, err := json.Marshal(sysSet)
	if err == nil {
		Write(string(marshal), settingFileName)
		SettingMapCache = sysSet
	}
}

// SystemSetting /**
type SystemSetting struct {
	// 系统语言设置
	Language string `json:"language"`

	// 剪切板历史相关的设置
	CopySetting CopySetting `json:"copySetting"`
}

// CopySetting /**
type CopySetting struct {
	// 剪切板历史最大条数
	MaxCount json.Number `json:"maxCount"`
	// 每条数据展示的最大条数是多少, 超过的部分使用 ...省略
	MaxCountOneData json.Number `json:"maxCountOneData"`
}
