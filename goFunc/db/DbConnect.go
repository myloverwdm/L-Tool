package db

// DatabaseOperation 数据库相关的操作
type DatabaseOperation interface {

	// TESTConnector 测试连接
	TESTConnector(dbInfo DataBaseInfo) bool
}
