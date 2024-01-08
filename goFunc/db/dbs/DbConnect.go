package dbs

import (
	"LTool/goFunc/db/info"
	"LTool/goFunc/global"
)

// DatabaseOperation 数据库相关的操作
type DatabaseOperation interface {

	// TESTConnector 测试连接
	TESTConnector(dbInfo info.DataBaseInfo) bool

	// GetDataBaseListByInfo 获取所有的库
	GetDataBaseListByInfo(dbInfo info.DataBaseInfo) global.LToolResponse
}
