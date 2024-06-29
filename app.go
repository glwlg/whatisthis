package main

import (
	"context"
	"fmt"
	"github.com/glwlg/whatisthis/internal/config"
	"github.com/glwlg/whatisthis/internal/gpt"
	"github.com/glwlg/whatisthis/internal/storage/nosql"
	"github.com/glwlg/whatisthis/internal/utils"
	"github.com/glwlg/whatisthis/pkg/robot"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx     context.Context
	cfg     config.Config
	storage nosql.Storage
}

var resultText string
var gptClient *gpt.Gpt

// NewApp creates a new App application struct
func NewApp() *App {

	storage, err := nosql.NewStorage()
	if err != nil {
		utils.LogError("无法初始化配置: %v", err)
	}
	// 初始化配置
	cfg, err := config.LoadConfig()
	if err != nil {
		utils.LogError("无法加载配置文件: %v", err)
	}

	a := &App{
		cfg:     *cfg,
		storage: *storage,
	}

	//utils.LogInfo(cfg.OpenAI.BaseUrl)

	//// 初始化全局热键和划词检测
	err = robot.InitGlobalHotkeys()
	if err != nil {
		utils.LogError("无法初始化全局热键:", err)
	}

	return a
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.hideWindow()

	guiConfig, err := a.storage.GetGuiConfig()
	if err != nil {
		utils.LogError("获取Gui配置失败: %v", err)
	}

	if !guiConfig.IsOpenAIEmpty() {
		// 初始化GPT模块
		gptClient = gpt.NewClient(a.cfg, guiConfig)
	}

	//// 注册划词检测回调
	robot.RegisterSelectionCallback(func(selectedText string) {
		utils.LogInfo("注册划词检测回调")

		if selectedText == lastSearchText || utils.IsBlankOrSpaces(selectedText) {
			utils.LogInfo("未选择文本")
			return
		}

		result, err := gptClient.Search(selectedText)
		if err != nil {
			utils.LogError("搜索出错: ", err)
			return
		}

		resultText = result

		runtime.EventsEmit(a.ctx, "onSearchResult", resultText)

		a.showWindow()

	})
}

// Greet returns a greeting for the given name
func (a *App) Greet() string {
	return resultText
}

func (a *App) hideWindow() {
	utils.LogInfo("hideWindow")
	runtime.Hide(a.ctx)
}

func (a *App) showWindow() {
	runtime.Show(a.ctx)
}

func (a *App) SetConfig(guiConfig *config.GuiConfig) {
	utils.LogInfo("SetConfig")
	fmt.Println(guiConfig)

	if !guiConfig.IsOpenAIEmpty() {
		err := a.storage.SaveOpenai(guiConfig.OpenAI)
		if err != nil {
			utils.LogError("保存配置失败: %v", err)
		}
		// 初始化GPT模块
		gptClient = gpt.NewClient(a.cfg, guiConfig)
	}

}

func (a *App) GetConfig() *config.GuiConfig {
	utils.LogInfo("GetConfig")

	guiConfig, err := a.storage.GetGuiConfig()
	if err != nil {
		utils.LogError("获取Gui配置失败: %v", err)
	}
	a.hideWindow()
	return guiConfig
}

//// Greet returns a greeting for the given name
//func (a *App) Greet(name string) string {
//	return fmt.Sprintf("Hello %s, It's show time!", name)
//}
