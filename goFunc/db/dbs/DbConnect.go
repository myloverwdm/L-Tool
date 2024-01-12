package dbs

import (
	"LTool/goFunc/db/info"
	"LTool/goFunc/global"
)

// DatabaseOperation 数据库相关的操作  regInfoName 注册信息的名称
type DatabaseOperation interface {

	// TESTConnector 测试连接
	TESTConnector(dbInfo info.DataBaseInfo) bool

	// GetDataBaseListByInfo 获取所有的库
	GetDataBaseListByInfo(dbInfo info.DataBaseInfo) global.LToolResponse

	// GetAllCharset 获取所有字符集和排序规则列表
	GetAllCharset(regInfoName string) map[string][]string

	// CreateDatabase 新建数据库
	CreateDatabase(regInfoName string, dataBaseCreateForm info.DataBaseCreateForm) global.LToolResponse

	// GetDataBaseDDL 得到某个dbName数据库的DDL语句
	GetDataBaseDDL(regInfoName, dbName string) string
}
