package dbStrategy

import "LTool/goFunc/db"

// MySqlStrategy 定义MySQL策略
type MySqlStrategy struct {
}

func (r MySqlStrategy) TESTConnector(dbInfo db.DataBaseInfo) bool {
	return true
}
