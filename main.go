package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

var lastSearchText = ""

func main() {

	// Create an instance of the app structure
	app := NewApp()
	//
	//AppMenu := menu.NewMenu()
	//
	//FileMenu := AppMenu.AddSubmenu("File")
	//robotMenu := FileMenu.AddText(robot.MenuTitle(), nil, func(c *menu.CallbackData) {
	//	robot.Toggle()
	//	runtime.EventsEmit(app.ctx, "robotMenuClick")
	//})
	//FileMenu.AddText("修改配置", nil, func(_ *menu.CallbackData) {
	//	runtime.EventsEmit(app.ctx, "openSetting")
	//})
	//FileMenu.AddText("退出", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
	//	runtime.Quit(app.ctx)
	//})

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
		AlwaysOnTop:       true,
		HideWindowOnClose: true,
		StartHidden:       true,
		OnDomReady: func(ctx context.Context) {
			//runtime.EventsOn(app.ctx, "robotMenuClick", func(optionalData ...interface{}) {
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
}
