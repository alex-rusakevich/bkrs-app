package bkrs

import "C"
import "unsafe"

// #cgo LDFLAGS: -luxtheme
// #include "gui.h"
import "C"

func MoveAppWindow(x int32, y int32) {
	C.moveAppWindow(C.int(x), C.int(y))
}

func MakeWindowAlwaysOnTop() {
	C.makeWindowAlwaysOnTop()
}

func LoadIcon() {
	C.loadIcon()
}

func MinimizeOnShortcut() {
	C.minimizeOnShortcut()
}

func InitGUI(window unsafe.Pointer) {
	C.initGUI(C.HWND(window))
}
