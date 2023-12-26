package onlineTools

type OnlineTool struct {
	Name string `json:"name"`

	Desc string `json:"desc"`

	Router string `json:"router"`
}

// 返回工具列表
func GetOnlineToolsList() []OnlineTool {

	return []OnlineTool{
		{"tool.json", "tool.json-desc", "/onlineTool/json"},
		{"tool.timestamp", "tool.timestamp-desc", "/onlineTool/timestamp"},
		{"tool.zk", "tool.zk-desc", "/onlineTool/zk"},
		{"tool.kafka", "tool.kafka-desc", "/onlineTool/kafka"},
	}
}
