package util

import (
	"LTool/goFunc"
	"LTool/goFunc/db/dbStrategy"
	"LTool/goFunc/db/info"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

func PingPort(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", address, time.Second*2)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func PingDb(databaseInfo info.DataBaseInfo) bool {
	port, err := strconv.Atoi(databaseInfo.Port)
	if err != nil {
		return false
	}
	if !PingPort(databaseInfo.Address, port) {
		return false
	}
	dbType := databaseInfo.DbType
	if dbType == "MySQL" {
		mysql := dbStrategy.MySqlStrategy{}
		return mysql.TESTConnector(databaseInfo)
	}
	return true
}

const dbFileCache = "cache/daCache"

func PingDbWithName(name string) bool {
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
			return PingDb(dataBaseInfoList[i])
		}
	}
	return false
}
