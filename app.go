package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/energye/systray"
	"github.com/glwlg/whatisthis/internal/config"
	"github.com/glwlg/whatisthis/internal/gpt"
	"github.com/glwlg/whatisthis/internal/storage/nosql"
	"github.com/glwlg/whatisthis/internal/utils"
	"github.com/glwlg/whatisthis/pkg/robot"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"os"
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
	config.InitConfig()
	cfg, err := config.LoadConfig()
	if err != nil {
		utils.LogError("无法加载配置文件: %v", err)
	}

	a := &App{
		cfg:     *cfg,
		storage: *storage,
	}

	return a
}

func (a *App) systemTray() {
	file, err := os.Open("./icon.ico")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建一个缓冲区读取器
	reader := bufio.NewReader(file)

	// 读取整个文件到 byte 数组
	icoByte, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	systray.SetIcon(icoByte) // read the icon from a file

	show := systray.AddMenuItem("打开", "打开面板")
	toggle := systray.AddMenuItem(robot.MenuTitle(), "开启/关闭划词")
	settings := systray.AddMenuItem("修改配置", "修改配置")
	apiConfig := systray.AddMenuItem("API配置", "API配置")
	systray.AddSeparator()
	exit := systray.AddMenuItem("退出", "退出")

	show.Click(func() { runtime.WindowShow(a.ctx) })
	toggle.Click(func() {
		a.RobotToggle(true)
	})
	settings.Click(func() {
		runtime.WindowShow(a.ctx)
		runtime.EventsEmit(a.ctx, "openSetting")
	})
	apiConfig.Click(func() {
		runtime.WindowShow(a.ctx)
		runtime.EventsEmit(a.ctx, "openApiConfig")
	})
	exit.Click(func() { os.Exit(0) })

	runtime.EventsOn(a.ctx, "robotToggle", func(optionalData ...interface{}) {
		toggle.SetTitle(robot.MenuTitle())
	})

	systray.SetOnClick(func(menu systray.IMenu) { runtime.WindowShow(a.ctx) })
	systray.SetOnRClick(func(menu systray.IMenu) {
		err := menu.ShowMenu()
		if err != nil {
			utils.LogError("右键菜单失败: %v", err)
			return
		}
	})
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	utils.LogInfo("开始初始化")

	guiConfig, err := a.storage.GetGuiConfig()
	if err != nil {
		utils.LogError("获取Gui配置失败: %v", err)
	}

	if !guiConfig.IsOpenAIEmpty() {
		// 初始化GPT模块
		gptClient = gpt.NewClient(a.cfg, guiConfig)
	}

	//// 初始化全局热键和划词检测
	err = robot.InitGlobalHotkeys()
	if err != nil {
		utils.LogError("无法初始化全局热键:", err)
	}

	//// 注册划词检测回调
	robot.RegisterSelectionCallback(func(selectedText string) {

		if selectedText == lastSearchText || utils.IsBlankOrSpaces(selectedText) {
			utils.LogInfo("未选择文本")
			return
		}
		lastSearchText = selectedText

		result, err := gptClient.Search(selectedText)
		if err != nil {
			utils.LogError("搜索出错: ", err)
			return
		}

		resultText = result

		runtime.EventsEmit(a.ctx, "onSearchResult", resultText)

		a.showWindow()

	})
	robot.RegisterCopyTextCallback(func() (string, error) {
		return runtime.ClipboardGetText(a.ctx)
	})

	utils.LogInfo("初始化完成1")

	//异步初始化systray

	systray.Run(a.systemTray, a.onExit)
	utils.LogInfo("初始化完成2")
}

func (a *App) onExit() {
	// clean up here
	runtime.Quit(a.ctx)
}

func (a *App) hideWindow() {
	utils.LogInfo("hideWindow")
	//runtime.Hide(a.ctx)
	runtime.WindowMinimise(a.ctx)
}

func (a *App) showWindow() {
	runtime.WindowUnminimise(a.ctx)
	runtime.Show(a.ctx)
}

func (a *App) RobotToggle(sendEvent bool) {
	robot.Toggle()
	if sendEvent {
		runtime.EventsEmit(a.ctx, "robotToggle")
	}
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
	if guiConfig.Delay != "" {
		err := a.storage.SaveConfig(guiConfig)
		if err != nil {
			utils.LogError("保存配置失败: %v", err)
		}
	}
}

func (a *App) GetConfig() *config.GuiConfig {
	utils.LogInfo("GetConfig")

	guiConfig, err := a.storage.GetGuiConfig()
	if err != nil {
		utils.LogError("获取Gui配置失败: %v", err)
	}
	return guiConfig
}

func (a *App) getClipboardText() string {
	clipboardText, err := runtime.ClipboardGetText(a.ctx)
	if err != nil {
		// 处理获取文本出错的情况，比如打印错误日志等
		runtime.LogError(a.ctx, "获取剪贴板文本出错"+err.Error())
		return ""
	}
	return clipboardText
}

func (a *App) RobotEnable() bool {
	enable := robot.Enable()
	return enable
}

//// Greet returns a greeting for the given name
//func (a *App) Greet(name string) string {
//	return fmt.Sprintf("Hello %s, It's show time!", name)
//}
