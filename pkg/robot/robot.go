package robot

import (
	"fmt"
	"github.com/glwlg/whatisthis/internal/utils"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"math"
	"time"
)

var mouseDrag = false
var enable = false

var selectionCallback func(string)
var copyTextCallback func() (string, error)

// InitGlobalHotkeys 初始化全局热键和划词检测
func InitGlobalHotkeys() error {
	utils.LogInfo("InitGlobalHotkey")

	// 开启一个协程用于监控鼠标事件
	go monitorMouseEvents()

	return nil
}

// RegisterSelectionCallback 注册划词检测回调函数
func RegisterSelectionCallback(callback func(string)) {
	utils.LogInfo("注册划词检测回调")
	selectionCallback = callback
}

// RegisterCopyTextCallback 注册复制文本回调函数
func RegisterCopyTextCallback(callback func() (string, error)) {
	utils.LogInfo("注册复制文本回调函数")
	copyTextCallback = callback
}

func Toggle() bool {
	enable = !enable
	utils.LogInfo(fmt.Sprintf("搜索状态：%v", enable))
	return enable
}

func Enable() bool {
	return enable
}

func MenuTitle() string {
	title := "开启划词"
	if enable {
		title = "关闭划词"
	}
	return title
}

// monitorMouseEvents 监控鼠标事件以检测划词操作
func monitorMouseEvents() {
	utils.LogInfo("monitorMouseEvents")

	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		if !enable {
			return
		}
		utils.LogInfo("MouseDown", enable, e)

		if e.Button != 1 {
			return
		}
		utils.LogInfo(e.String())
		if (e.Clicks == 1 || e.Clicks == 0) && mouseDrag {
			mouseDrag = false
			executeCallback(e)
		} else if e.Clicks == 2 {
			executeCallback(e)
		}

	})

	hook.Register(hook.MouseDrag, []string{}, func(e hook.Event) {
		//utils.LogInfo("MouseDrag")
		if !enable {
			return
		}
		mouseDrag = true
	})

	utils.LogInfo("hook.Start")
	s := hook.Start()
	<-hook.Process(s)
	utils.LogInfo("hook.Process")
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
	err := robotgo.KeyTap("insert", "ctrl")
	if err != nil {
		utils.LogError("copy text Error", err)
		return "", err
	}

	clipboardText, err := copyTextCallback()
	if err != nil {
		utils.LogError("clipboard  ReadAll Error", err)
		return "", err
	}
	return clipboardText, nil
}
