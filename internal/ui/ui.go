package ui

//import (
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/app"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/widget"
//	"github.com/glwlg/whatisthis/internal/config"
//	"github.com/glwlg/whatisthis/internal/gpt"
//	"github.com/glwlg/whatisthis/internal/utils"
//	"github.com/glwlg/whatisthis/pkg/robot"
//	"log"
//)
//
//type App struct {
//	Window  fyne.Window
//	search  *widget.Entry
//	results *widget.List
//	//gptAPI         *gpt.Client
//}
//
//var lastSearchText = ""
//
//func NewApp(cfg *config.Config) (fyne.App, error) {
//	a := app.New()
//	w := a.NewWindow("这是什么？")
//
//	icon, _ := fyne.LoadResourceFromPath("golang.ico")
//	a.SetIcon(icon)
//
//	w.CenterOnScreen()
//	w.Hide()
//
//	// 初始化GPT模块
//	gptClient := gpt.NewClient(cfg)
//
//	//// 初始化历史记录模块
//	//historyManager := history.NewManager()
//	//
//	//// 初始化收藏模块
//	//favoritesManager := favorites.NewManager()
//	//
//	//// 初始化推荐模块
//	//recommendationEngine := recommendation.NewEngine(historyManager)
//
//	// 初始化GUI
//	//uiManager, err := ui.NewApp()
//	//if err != nil {
//	//	log.Fatalf("无法初始化UI: %v", err)
//	//}
//	//
//	//// 初始化全局热键和划词检测
//	err := robot.InitGlobalHotkeys()
//	if err != nil {
//		log.Fatalf("无法初始化全局热键: %v", err)
//	}
//
//	//// 注册划词检测回调
//	robot.RegisterSelectionCallback(func(selectedText string) {
//
//		if selectedText == lastSearchText || utils.IsBlankOrSpaces(selectedText) {
//			utils.LogInfo("未选择文本")
//			return
//		}
//
//		result, err := gptClient.Search(selectedText)
//		if err != nil {
//			utils.LogError("搜索出错: ", err)
//			return
//		}
//
//		resultLabel := widget.NewLabel(result)
//		var copyButton *widget.Button
//		copyButton = widget.NewButton("copy", func() {
//			_ = robot.WriteToClipboard(result)
//
//			copyButton.SetText("copied!")
//		})
//
//		w.SetContent(container.NewVBox(
//			resultLabel,
//			copyButton,
//		))
//		w.Show()
//
//	})
//
//	return a, nil
//}
//
//func (a *App) TriggerSearch() {
//	// 可以在这里接收划词内容，然后调用搜索功能
//	a.search.SetText("划词内容")
//}
