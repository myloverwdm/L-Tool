package main

import (
	"LTool/configuration"
	"LTool/goFunc"
	"LTool/goFunc/index"
	"LTool/page/onlineTools"
	"context"
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

func (a *App) GetNowLanguage() string {
	return index.GetNowLanguage()
}

func (a *App) SetNowLanguage(language string) {
	index.SetNowLanguage(language)
}
