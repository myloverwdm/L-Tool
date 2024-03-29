package main

import (
	"LTool/configuration"
	"LTool/goFunc"
	"LTool/goFunc/db"
	"LTool/goFunc/db/funcDb"
	"LTool/goFunc/db/info"
	"LTool/goFunc/global"
	"LTool/goFunc/index"
	"LTool/goFunc/index/code"
	"LTool/goFunc/index/copy"
	"LTool/goFunc/util"
	"LTool/page/onlineTools"
	"context"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/pkg/browser"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
}

func (a *App) GetItemList() []configuration.Index {
	return configuration.IndexRoute
}

func (a *App) GetLanguageList() []configuration.Language {
	return configuration.LanguageList
}

func (a *App) GetOnlineToolsList() []onlineTools.OnlineTool {
	return onlineTools.GetOnlineToolsList()
}

func (a *App) ReadZkNode(path string) index.ZkData {
	return index.ReadZk(path)
}

func (a *App) InitZk() index.ZkForm {
	return index.InitZk()
}
func (a *App) SetForm(zkForm index.ZkForm) string {
	return index.SetForm(zkForm)
}

func (a *App) DeleteZk(path string) string {
	return index.DeleteZk(path)
}

func (a *App) AddPath(path, newPath, value string) string {
	return index.AddPath(path, newPath, value)
}

func (a *App) JsonToXml(value string) goFunc.LTResponse {
	return index.JsonToXml(value)
}

func (a *App) GetCommand() []index.CommandPojo {
	return index.GetCommand()
}

func (a *App) GetTopicAndPartition() []index.TopicAndPartition {
	return index.GetTopicAndPartition()
}

func (a *App) GetKafkaForm() index.KafkaForm {
	return index.GetKafkaForm()
}

func (a *App) UpdateKafkaForm(kafkaForm index.KafkaForm) string {
	return index.UpdateKafkaForm(kafkaForm)
}

func (a *App) BrowserOpenURL(url string) {
	err := browser.OpenURL(url)
	if err != nil {
		return
	}
}

var lastCopyData = ""

func (a *App) GetCopyData() {
	content, err := clipboard.ReadAll()
	if err != nil {
		return
	}
	if content == "" {
		return
	}
	if lastCopyData == content {
		return
	}
	lastCopyData = content
	copyHis := copy.GetCopyHis()
	if len(copyHis) > 0 && copyHis[len(copyHis)-1].Data == content {
		return
	}
	// 更新新的剪切板
	copyHis = append(copyHis, copy.CopyHis{Data: content, Time: time.Now().Format("2006-01-02 15:04")})

	allSettings := goFunc.GetOrDefaultAllSettings()
	i, err := allSettings.CopySetting.MaxCount.Int64()
	maxCount := int(i)
	if len(copyHis) > maxCount {
		copyHis = copyHis[len(copyHis)-maxCount:]
	}
	copy.SetCopyHis(copyHis)
}

func (a *App) GetCopyHis() []copy.CopyHis {
	return copy.GetCopyHis()
}

func (a *App) GetCodeGroup() []code.CodeGroup {
	return code.GetCodeGroup()
}

func (a *App) GetOrDefaultAllSettings() goFunc.SystemSetting {
	return goFunc.GetOrDefaultAllSettings()
}

func (a *App) UpdateSystemSettings(setting string) {
	goFunc.UpdateSystemSettings(setting)
}

func (a *App) UpdateFileCache(filename, content string) {
	goFunc.UpdateFileCache(filename, content)
}

func (a *App) GetFileCache(filename string) string {
	return goFunc.GetFileCache(filename)
}

func (a *App) GetAllDataBaseInfo() []info.DataBaseInfo {
	return db.GetAllDataBaseInfo()
}

func (a *App) AddOrUpdateDataBaseInfo(databaseInfo info.DataBaseInfo, isAdd bool) string {
	return db.AddOrUpdateDataBaseInfo(databaseInfo, isAdd)
}

func (a *App) DeleteOneDataBaseInfo(dbName string) string {
	return db.DeleteOneDataBaseInfo(dbName)
}

func (a *App) PingDb(databaseInfo info.DataBaseInfo) bool {
	return util.PingDb(databaseInfo)
}

func (a *App) PingDbWithName(name string) bool {
	return util.PingDbWithName(name)
}

func (a *App) GetDataBaseListByRegName(name string) global.LToolResponse {
	return db.GetDataBaseListByRegName(name)
}

func (a *App) GetDataBaseType(name string) string {
	return db.GetDataBaseType(name)
}

func (a *App) T1(name string) info.DataBaseCreateForm {
	fmt.Println("仅是为了把DataBaseCreateForm给到前端")
	return info.DataBaseCreateForm{}
}

// GetAllCharset 获取所有的字符集和排序规则
func (a *App) GetAllCharset(name, dbType string) map[string][]string {
	fmt.Println("仅是为了把DataBaseCreateForm给到前端")
	return db.GetAllCharset(name, dbType)
}

func (a *App) CreateDatabase(name, dbType string, dataBaseCreateForm info.DataBaseCreateForm) global.LToolResponse {
	return funcDb.CreateDatabase(dbType, name, dataBaseCreateForm)
}

func (a *App) GetDataBaseDDL(dbType, regInfoName, dbName string) string {
	return funcDb.GetDataBaseDDL(dbType, regInfoName, dbName)
}
