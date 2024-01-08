package dbStrategy

import (
	"LTool/goFunc/db/dbcache"
	"LTool/goFunc/db/info"
	"LTool/goFunc/global"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

// MySqlStrategy 定义MySQL策略
type MySqlStrategy struct {
}

func (r MySqlStrategy) TESTConnector(dbInfo info.DataBaseInfo) bool {

	return true
}

func (r MySqlStrategy) GetDataBaseListByInfo(dbInfo info.DataBaseInfo) global.LToolResponse {
	db := dbcache.DbMap[dbInfo.Name]
	if db == nil {
		dbcache.CreateConnection(dbInfo)
		db = dbcache.DbMap[dbInfo.Name]
		if db == nil {
			return global.LToolResponse{Message: "数据库连接失败"}
		}
	}
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		return global.LToolResponse{Message: "获取库列表失败"}
	}
	defer rows.Close()
	var databaseNames []string
	// 遍历查询结果，获取数据库名称
	for rows.Next() {
		var databaseName string
		err := rows.Scan(&databaseName)
		if err != nil {
			return global.LToolResponse{Message: "获取库列表失败"}
		}
		databaseNames = append(databaseNames, databaseName)
	}
	marshal, _ := json.Marshal(databaseNames)
	return global.LToolResponse{Success: true, Data: string(marshal)}
}
