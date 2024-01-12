package db

import (
	"LTool/goFunc"
	"LTool/goFunc/db/dbcache"
	"LTool/goFunc/db/funcDb"
	"LTool/goFunc/db/info"
	"LTool/goFunc/global"
	"encoding/json"
	"sort"
	"time"
)

func GetAllDataBaseInfo() []info.DataBaseInfo {
	fileCache := goFunc.GetFileCache(dbFileCache)
	var dataBaseInfoList []info.DataBaseInfo
	if fileCache == "" {
		return []info.DataBaseInfo{}
	} else {
		err2 := json.Unmarshal([]byte(fileCache), &dataBaseInfoList)
		if err2 != nil {
			dataBaseInfoList = []info.DataBaseInfo{}
		}
	}
	sort.Slice(dataBaseInfoList, func(i, j int) bool {
		return dataBaseInfoList[i].UpdateTime > dataBaseInfoList[j].UpdateTime
	})
	return dataBaseInfoList
}

const dbFileCache = "cache/daCache"

func AddOrUpdateDataBaseInfo(databaseInfo info.DataBaseInfo, isAdd bool) string {
	fileCache := goFunc.GetFileCache(dbFileCache)
	var dataBaseInfoList []info.DataBaseInfo
	if fileCache == "" {
		dataBaseInfoList = []info.DataBaseInfo{}
	} else {
		err2 := json.Unmarshal([]byte(fileCache), &dataBaseInfoList)
		if err2 != nil {
			dataBaseInfoList = []info.DataBaseInfo{}
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
	var dataBaseInfoList []info.DataBaseInfo
	if fileCache == "" {
		dataBaseInfoList = []info.DataBaseInfo{}
	} else {
		err2 := json.Unmarshal([]byte(fileCache), &dataBaseInfoList)
		if err2 != nil {
			dataBaseInfoList = []info.DataBaseInfo{}
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
	// 是否已经建立了连接,若是,则断开
	db := dbcache.DbMap[dbName]
	if db == nil {
		db.Close()
		delete(dbcache.DbMap, dbName)
	}
	return ""
}

func GetDataBaseListByRegName(name string) global.LToolResponse {
	fileCache := goFunc.GetFileCache(dbFileCache)
	var dataBaseInfoList []info.DataBaseInfo
	if fileCache == "" {
		dataBaseInfoList = []info.DataBaseInfo{}
	} else {
		err2 := json.Unmarshal([]byte(fileCache), &dataBaseInfoList)
		if err2 != nil {
			dataBaseInfoList = []info.DataBaseInfo{}
		}
	}
	for i := 0; i < len(dataBaseInfoList); i++ {
		if dataBaseInfoList[i].Name == name {
			return funcDb.GetDataBaseListByInfo(dataBaseInfoList[i])
		}
	}
	return global.LToolResponse{}
}

func GetDataBaseType(name string) string {
	fileCache := goFunc.GetFileCache(dbFileCache)
	var dataBaseInfoList []info.DataBaseInfo
	if fileCache == "" {
		dataBaseInfoList = []info.DataBaseInfo{}
	} else {
		err2 := json.Unmarshal([]byte(fileCache), &dataBaseInfoList)
		if err2 != nil {
			dataBaseInfoList = []info.DataBaseInfo{}
		}
	}
	for i := 0; i < len(dataBaseInfoList); i++ {
		if dataBaseInfoList[i].Name == name {
			return dataBaseInfoList[i].DbType
		}
	}
	return ""
}

func GetAllCharset(name, dbType string) map[string][]string {
	return funcDb.GetAllCharset(name, dbType)
}
