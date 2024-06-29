package hotkey

//
//import (
//	"github.com/glwlg/whatisthis/internal/utils"
//	"github.com/robotn/gorobotgo"
//)
//
//// HotkeyManager 管理全局热键
//type HotkeyManager struct {
//	hk gorobotgo.Hotkey
//}
//
//// NewHotkey 创建一个新的 HotkeyManager 实例
//func NewHotkey(hotkey string) (*HotkeyManager, error) {
//	h, err := gorobotgo.NewHotkey(hotkey)
//	if err != nil {
//		utils.LogError("Failed to create hotkey", err)
//		return nil, err
//	}
//	return &HotkeyManager{hk: h}, nil
//}
//
//// OnHotkey 设置热键触发的回调函数
//func (h *HotkeyManager) OnHotkey(callback func()) {
//	h.hk.OnHotkey(callback)
//}
//
//// Close 关闭热键监听器
//func (h *HotkeyManager) Close() {
//	h.hk.Close()
//}
