package info

type DataBaseInfo struct {
	// 数据库名称
	Name string `json:"name"`
	// 数据库的类型
	DbType string `json:"dbType"`
	// 地址
	Address string `json:"address"`
	// Port
	Port string `json:"port"`
	// 用户名
	UserName string `json:"userName"`
	// 密码
	Password string `json:"password"`

	// 注册时间
	RegTime int64 `json:"regTime"`
	// 更新时间
	UpdateTime int64 `json:"updateTime"`

	RegTimeStr    string `json:"RegTimeStr"`
	UpdateTimeStr string `json:"updateTimeStr"`
}

type DataBaseCreateForm struct {
	// 数据库名称
	DataBaseName string `json:"dataBaseName"`
	// 字符集
	CharacterSet string `json:"characterSet"`
	// 排序规则
	Collation string `json:"collation"`
	// 模式/用户名
	UserName string `json:"userName"`
	// 密码
	Password string `json:"password"`
}
