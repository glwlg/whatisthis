package main

import (
	"context"
	"embed"
	"github.com/glwlg/whatisthis/internal/utils"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"log"
	"os"
)

//go:embed all:frontend/dist
var assets embed.FS

var lastSearchText = ""

func main() {

	// 创建一个日志文件
	logF, e := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		log.Fatal(e)
	}
	defer logF.Close()
	// 设置日志输出到文件
	log.SetOutput(logF)

	// Create an instance of the app structure
	app := NewApp()
	//
	//AppMenu := menu.NewMenu()
	//
	//FileMenu := AppMenu.AddSubmenu("File")
	//robotMenu := FileMenu.AddText(robot.MenuTitle(), nil, func(c *menu.CallbackData) {
	//	robot.Toggle()
	//	runtime.EventsEmit(app.ctx, "robotToggle")
	//})
	//FileMenu.AddText("修改配置", nil, func(_ *menu.CallbackData) {
	//	runtime.EventsEmit(app.ctx, "openSetting")
	//})
	//FileMenu.AddText("退出", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
	//	runtime.Quit(app.ctx)
	//})
	utils.LogInfo("初始化 app")

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "whatisthis",
		Width:  500,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour:  &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:         app.startup,
		AlwaysOnTop:       false,
		HideWindowOnClose: true,
		StartHidden:       false,
		OnDomReady: func(ctx context.Context) {
			//runtime.EventsOn(app.ctx, "robotToggle", func(optionalData ...interface{}) {
			//	robotMenu.SetLabel(robot.MenuTitle())
			//})
		},
		Frameless: true,
		//OnBeforeClose: func(ctx context.Context) (prevent bool) {
		//	return true
		//},

		//Menu: AppMenu,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

	utils.LogInfo("初始化app完成")
	//systray.Run(app.systemTray, func() {})
	utils.LogInfo("托盘初始化完成")

}
