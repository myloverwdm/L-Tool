package funcDb

import (
	"LTool/goFunc/db/dbStrategy"
	"LTool/goFunc/db/info"
	"LTool/goFunc/global"
)

func GetDataBaseListByInfo(dataBaseInfo info.DataBaseInfo) global.LToolResponse {
	dbType := dataBaseInfo.DbType
	if dbType == "MySQL" {
		mysql := dbStrategy.MySqlStrategy{}
		return mysql.GetDataBaseListByInfo(dataBaseInfo)
	}
	return global.LToolResponse{}
}
