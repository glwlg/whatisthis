package robot

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/glwlg/whatisthis/internal/utils"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"math"
	"time"
)

var selectionCallback func(string)

// InitGlobalHotkeys 初始化全局热键和划词检测
func InitGlobalHotkeys() error {
	// 可以在这里设置一些全局热键，如果需要的话
	// 例如：robot.RegisterHotKey("ctrl", "alt", "c", func() { log.Println("Hotkey pressed!") })

	// 开启一个协程用于监控鼠标事件
	go monitorMouseEvents()

	return nil
}

// RegisterSelectionCallback 注册划词检测回调函数
func RegisterSelectionCallback(callback func(string)) {
	selectionCallback = callback
}

var mouseDrag = false
var enable = false

// monitorMouseEvents 监控鼠标事件以检测划词操作
func monitorMouseEvents() {
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		if !enable {
			return
		}
		if e.Button != 1 {
			return
		}
		utils.LogInfo("MouseDown")
		utils.LogInfo(e.String())
		if (e.Clicks == 1 || e.Clicks == 0) && mouseDrag {
			mouseDrag = false
			executeCallback(e)
		} else if e.Clicks == 2 {
			executeCallback(e)
		}

	})

	hook.Register(hook.MouseDrag, []string{}, func(e hook.Event) {
		if !enable {
			return
		}
		mouseDrag = true
	})

	//hook.Register(hook.KeyDown, []string{"q", "ctrl", "alt"}, func(e hook.Event) {
	//	enable = !enable
	//	utils.LogInfo(fmt.Sprintf("搜索状态：%v", enable))
	//})

	hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
		utils.LogInfo(fmt.Sprintf("e：%v", e))
	})
	//hook.Register(hook.MouseMove, []string{}, func(e hook.Event) {
	//	//utils.LogInfo("MouseMove")
	//	executed = true
	//})
	//
	//hook.Register(hook.MouseUp, []string{}, func(e hook.Event) {
	//	utils.LogInfo("MouseUp")
	//	executed = true
	//	//executeCallback()
	//
	//	//hook.End()
	//})

	s := hook.Start()
	<-hook.Process(s)
	//
	//for {
	//	mleft := hook.AddEvent("mleft")
	//	if mleft {
	//		utils.LogInfo("Mouse left button pressed")
	//	}
	//
	//	//err := robot.KeyTap("c", "ctrl")
	//	//if err != nil {
	//	//	utils.LogError("copy text Error", err)
	//	//	return
	//	//}
	//	//
	//	//clipboardText, err := clipboard.ReadAll()
	//	//if err != nil {
	//	//	utils.LogError("clipboard  ReadAll Error", err)
	//	//	return
	//	//}
	//	//selectionCallback(clipboardText)
	//	//
	//	//	// 获取选中的文本
	//	//selectedText := robot.GetSelection()
	//	//log.Printf("Selected text: %s", selectedText)
	//	//
	//	//	// 触发回调函数
	//	//if selectionCallback != nil && selectedText != "" {
	//	//	selectionCallback(selectedText)
	//	//}
	//	//}
	//	time.Sleep(100 * time.Millisecond) // 防止CPU占用过高
	//}
}

// monitorMouseEvents 监控鼠标事件以检测划词操作
func executeCallback(e hook.Event) {
	eX := int(e.X)
	eY := int(e.Y)

	utils.LogInfo(fmt.Sprintf("eX:%d, eY:%d", eX, eY))

	x, y := robotgo.Location()
	utils.LogInfo(fmt.Sprintf("x:%d, y:%d", x, y))
	time.Sleep(1000 * time.Millisecond)
	newX, newY := robotgo.Location()
	utils.LogInfo(fmt.Sprintf("newX:%d, newY:%d", newX, newY))

	// 计算x和y坐标的变动绝对值
	deltaX := math.Abs(float64(newX - x))
	deltaY := math.Abs(float64(newY - y))

	// 判断变动绝对值是否小于50
	if deltaX < 50 && deltaY < 50 {
		utils.LogInfo("坐标变动绝对值小于10")
	} else {
		utils.LogInfo("坐标变动绝对值大于或等于10")
		return
	}

	text, _ := CopyText()

	if len(text) > 0 {
		selectionCallback(text)
	}

}

func CopyText() (string, error) {
	err := robotgo.KeyTap("c", "ctrl")
	if err != nil {
		utils.LogError("copy text Error", err)
		return "", err
	}

	clipboardText, err := clipboard.ReadAll()
	if err != nil {
		utils.LogError("clipboard  ReadAll Error", err)
		return "", err
	}
	return clipboardText, nil
}

func WriteToClipboard(text string) error {
	err := clipboard.WriteAll(text)
	if err != nil {
		utils.LogError("clipboard  WriteAll Error", err)
		return err
	}
	return nil
}
