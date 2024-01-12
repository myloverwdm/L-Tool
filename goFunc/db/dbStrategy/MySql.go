package dbStrategy

import (
	"LTool/goFunc/db/dbcache"
	"LTool/goFunc/db/info"
	"LTool/goFunc/global"
	"encoding/json"
	"fmt"
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

// GetAllCharset 获取Mysql所有字符集和排序规则列表
func (r MySqlStrategy) GetAllCharset(daName string) map[string][]string {
	db := dbcache.DbMap[daName]
	var result = make(map[string][]string)
	if db == nil {
		return result
	}
	// 执行SQL查询语句，获取所有数据库名称
	rows, err := db.Query("SHOW COLLATION")
	if err != nil {
		return result
	}
	defer rows.Close()
	// 遍历查询结果，获取数据库名称
	for rows.Next() {
		var collation string
		var charset string
		var temp1, temp2, temp3, temp4 interface{} // 临时变量用于接收不需要的列值
		err := rows.Scan(&collation, &charset, &temp1, &temp2, &temp3, &temp4)
		if err != nil {
			return result
		}
		result[charset] = append(result[charset], collation)
	}
	return result
}

// CreateDatabase 新建数据库
func (r MySqlStrategy) CreateDatabase(regInfoName string, dataBaseCreateForm info.DataBaseCreateForm) global.LToolResponse {
	db := dbcache.DbMap[regInfoName]
	if db == nil {
		return global.LToolResponse{Message: "数据库连接失败"}
	}
	SQL := "CREATE DATABASE " + dataBaseCreateForm.DataBaseName + " CHARACTER SET " + dataBaseCreateForm.CharacterSet + " COLLATE " + dataBaseCreateForm.Collation
	_, err := db.Exec(SQL)
	if err != nil {
		fmt.Println(err.Error())
		return global.LToolResponse{Message: "创建库失败"}
	}
	return global.LToolResponse{Success: true, Message: "创建库成功"}
}

func (r MySqlStrategy) GetDataBaseDDL(regInfoName, dbName string) string {
	db := dbcache.DbMap[regInfoName]
	if db == nil {
		return ""
	}
	SQL := "SHOW CREATE DATABASE " + dbName
	rows, err := db.Query(SQL)
	if err != nil {
		return ""
	}
	defer rows.Close()
	// 遍历查询结果，获取数据库名称
	for rows.Next() {
		var ddl string
		var temp string
		err := rows.Scan(&temp, &ddl)
		if err != nil {
			return ""
		}
		return ddl
	}
	return ""
}
