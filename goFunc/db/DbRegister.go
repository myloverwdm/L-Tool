package db

import (
	"LTool/goFunc"
	"encoding/json"
	"sort"
	"time"
)

type DataBaseInfo struct {
	// 数据库名称
	Name string `json:"name"`
	// 数据库的类型
	DbType string `json:"dbType"`
	// 地址
	Address string `json:"address"`
	// Port
	Port string `json:"port"`
	// 用户名
	UserName string `json:"userName"`
	// 密码
	Password string `json:"password"`

	// 注册时间
	RegTime int64 `json:"regTime"`
	// 更新时间
	UpdateTime int64 `json:"updateTime"`

	RegTimeStr    string `json:"RegTimeStr"`
	UpdateTimeStr string `json:"updateTimeStr"`
}

func GetAllDataBaseInfo() []DataBaseInfo {
	fileCache := goFunc.GetFileCache(dbFileCache)
	var dataBaseInfoList []DataBaseInfo
	if fileCache == "" {
		return []DataBaseInfo{}
	} else {
		err2 := json.Unmarshal([]byte(fileCache), &dataBaseInfoList)
		if err2 != nil {
			dataBaseInfoList = []DataBaseInfo{}
		}
	}
	sort.Slice(dataBaseInfoList, func(i, j int) bool {
		return dataBaseInfoList[i].UpdateTime > dataBaseInfoList[j].UpdateTime
	})
	return dataBaseInfoList
}

const dbFileCache = "cache/daCache"

func AddOrUpdateDataBaseInfo(databaseInfo DataBaseInfo, isAdd bool) string {
	fileCache := goFunc.GetFileCache(dbFileCache)
	var dataBaseInfoList []DataBaseInfo
	if fileCache == "" {
		dataBaseInfoList = []DataBaseInfo{}
	} else {
		err2 := json.Unmarshal([]byte(fileCache), &dataBaseInfoList)
		if err2 != nil {
			dataBaseInfoList = []DataBaseInfo{}
		}
	}
	name := databaseInfo.Name
	already := false
	index := 0
	for i := 0; i < len(dataBaseInfoList); i++ {
		dataBase := dataBaseInfoList[i]
		if name == dataBase.Name {
			index = i
			already = true
			break
		}
	}
	if already && isAdd {
		return "db.same-db"
	}
	if !already && !isAdd {
		// 更新的, 名称不可更改
		return "db.db-404"
	}
	if isAdd {
		now := time.Now()
		databaseInfo.UpdateTime = now.UnixMilli()
		databaseInfo.RegTime = now.UnixMilli()
		dataBaseInfoList = append(dataBaseInfoList, databaseInfo)
	} else {
		// 是更新
		databaseInfo.UpdateTime = time.Now().UnixMilli()
		dataBaseInfoList = append(append(dataBaseInfoList[:index], dataBaseInfoList[index+1:]...), databaseInfo)
	}
	marshal, err := json.Marshal(dataBaseInfoList)
	if err != nil {
		return "globals.unknown-error"
	}
	goFunc.UpdateFileCache(dbFileCache, string(marshal))
	return ""
}

func DeleteOneDataBaseInfo(dbName string) string {
	fileCache := goFunc.GetFileCache(dbFileCache)
	var dataBaseInfoList []DataBaseInfo
	if fileCache == "" {
		dataBaseInfoList = []DataBaseInfo{}
	} else {
		err2 := json.Unmarshal([]byte(fileCache), &dataBaseInfoList)
		if err2 != nil {
			dataBaseInfoList = []DataBaseInfo{}
		}
	}
	index := -1
	for i := 0; i < len(dataBaseInfoList); i++ {
		dataBase := dataBaseInfoList[i]
		if dbName == dataBase.Name {
			index = i
			break
		}
	}
	if index < 0 {
		return "db.db-404"
	}
	dataBaseInfoList = append(dataBaseInfoList[:index], dataBaseInfoList[index+1:]...)
	marshal, err := json.Marshal(dataBaseInfoList)
	if err != nil {
		return "globals.unknown-error"
	}
	goFunc.UpdateFileCache(dbFileCache, string(marshal))
	return ""
}
