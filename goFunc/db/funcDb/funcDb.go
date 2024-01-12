package funcDb

import (
	"LTool/goFunc/db/dbStrategy"
	"LTool/goFunc/db/dbs"
	"LTool/goFunc/db/info"
	"LTool/goFunc/global"
)

func GetDbStrategy(dbType string) dbs.DatabaseOperation {
	var database dbs.DatabaseOperation
	if dbType == "MySQL" {
		database = dbStrategy.MySqlStrategy{}
	}
	return database
}

func GetDataBaseListByInfo(dataBaseInfo info.DataBaseInfo) global.LToolResponse {
	dbType := dataBaseInfo.DbType
	return GetDbStrategy(dbType).GetDataBaseListByInfo(dataBaseInfo)
}

func GetAllCharset(name, dbType string) map[string][]string {
	return GetDbStrategy(dbType).GetAllCharset(name)
}

func CreateDatabase(dbType, dbName string, dataBaseCreateForm info.DataBaseCreateForm) global.LToolResponse {
	return GetDbStrategy(dbType).CreateDatabase(dbName, dataBaseCreateForm)
}

func GetDataBaseDDL(dbType, regInfoName, dbName string) string {
	return GetDbStrategy(dbType).GetDataBaseDDL(regInfoName, dbName)
}
