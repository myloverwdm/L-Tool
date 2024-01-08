package dbcache

import (
	"LTool/goFunc/db/info"
	"database/sql"
	"fmt"
)

var DbMap = map[string]*sql.DB{}

func CreateConnection(info info.DataBaseInfo) {
	// 连接字符串，包括数据库用户名、密码、主机和端口
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/", info.UserName, info.Password, info.Address, info.Port)
	// 创建数据库连接
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("无法连接到数据库：", err)
		return
	}
	// 测试连接是否成功
	err = db.Ping()
	if err != nil {
		fmt.Println("无法连接到数据库2：", err)
		return
	}
	DbMap[info.Name] = db
	return
}
