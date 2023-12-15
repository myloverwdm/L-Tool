package index

import (
	"LTool/goFunc"
	"encoding/json"
	"github.com/go-zookeeper/zk"
	"time"
)

type ZkData struct {
	NextPath []string `json:"nextPath"`
	Data     string   `json:"data"`
	StatData Stat     `json:"statData"`
}

type ZkForm struct {
	Address string `json:"address"`
	ZkData  ZkData `json:"zkData"`
}

type Stat struct {
	Czxid          int64 `json:"czxid,omitempty"`          // The zxid of the change that caused this znode to be created.
	Mzxid          int64 `json:"mzxid,omitempty"`          // The zxid of the change that last modified this znode.
	Ctime          int64 `json:"ctime,omitempty"`          // The time in milliseconds from epoch when this znode was created.
	Mtime          int64 `json:"mtime,omitempty"`          // The time in milliseconds from epoch when this znode was last modified.
	Version        int32 `json:"version,omitempty"`        // The number of changes to the data of this znode.
	Cversion       int32 `json:"cversion,omitempty"`       // The number of changes to the children of this znode.
	Aversion       int32 `json:"aversion,omitempty"`       // The number of changes to the ACL of this znode.
	EphemeralOwner int64 `json:"ephemeralOwner,omitempty"` // The session id of the owner of this znode if the znode is an ephemeral node. If it is not an ephemeral node, it will be zero.
	DataLength     int32 `json:"dataLength,omitempty"`     // The length of the data field of this znode.
	NumChildren    int32 `json:"numChildren,omitempty"`    // The number of children of this znode.
	Pzxid          int64 `json:"pzxid,omitempty"`          // last modified children
}

func ReadZk(path string) ZkData {
	// 连接到ZooKeeper服务器
	var zkForm ZkForm
	zkForm = goFunc.Get(zookeeperKey, zkForm)
	conn, _, err := zk.Connect([]string{zkForm.Address}, time.Second*30)
	if err != nil {
		// fmt.Println("连接失败", err)
		return ZkData{}
	}
	defer closeZkConnection(conn)
	// 读取节点数据
	nodeData, stat, err := conn.Get(path)
	if err != nil {
		return ZkData{}
	}
	// fmt.Println(nodeData, stat)
	var data = ""
	if nodeData != nil {
		data = string(nodeData)
	}
	children, _, err := conn.Children(path)
	if err != nil {
		return ZkData{}
	}
	return ZkData{
		NextPath: children,
		Data:     data,
		StatData: Stat{
			Czxid:          stat.Czxid,
			Mzxid:          stat.Mzxid,
			Ctime:          stat.Ctime,
			Mtime:          stat.Mtime,
			Version:        stat.Version,
			Cversion:       stat.Cversion,
			Aversion:       stat.Aversion,
			EphemeralOwner: stat.EphemeralOwner,
			DataLength:     stat.DataLength,
			NumChildren:    stat.NumChildren,
			Pzxid:          stat.Pzxid,
		},
	}
}

func DeleteZk(path string) string {
	var zkForm ZkForm
	zkForm = goFunc.Get(zookeeperKey, zkForm)
	conn, _, err := zk.Connect([]string{zkForm.Address}, time.Second*30)
	if err != nil {
		// fmt.Println("连接失败", err)
		return err.Error()
	}
	defer closeZkConnection(conn)
	// fmt.Println("删除的地址", zkForm.Address, " 删除的路径", path)
	err1 := conn.Delete(path, -1)
	if err1 != nil {
		// 错误处理
		return err1.Error()
	}
	return ""
}

func AddPath(path, newPath, value string) string {
	var zkForm ZkForm
	zkForm = goFunc.Get(zookeeperKey, zkForm)
	conn, _, err := zk.Connect([]string{zkForm.Address}, time.Second*30)
	if err != nil {
		// fmt.Println("连接失败", err)
		return err.Error()
	}
	defer closeZkConnection(conn)
	data := []byte(value)          // 节点的数据，以字节数组形式提供
	flags := int32(0)              // 节点的标志，通常为 0
	acl := zk.WorldACL(zk.PermAll) // 节点的 ACL（访问控制列表），这里使用全球可见 ACL // 节点的路径
	_, err2 := conn.Create(path+"/"+newPath, data, flags, acl)
	if err2 != nil {
		// 错误处理
		return err.Error()
	} else {
		return ""
	}

}

func closeZkConnection(conn *zk.Conn) {
	conn.Close()
}

const zookeeperKey = "Zookeeper"

func InitZk() ZkForm {
	var zkForm1 ZkForm
	var zkForm = goFunc.Get(zookeeperKey, zkForm1)
	if zkForm.Address != "" {
		zkForm.ZkData = ReadZk("/")
	}
	// 输出结构体的内容
	return zkForm
}

func SetForm(zkForm ZkForm) string {
	marshal, err := json.Marshal(zkForm)
	if err == nil {
		goFunc.Set(zookeeperKey, string(marshal))
		goFunc.GlobalMap[zookeeperKey] = string(marshal)
	}
	return "success"
}
