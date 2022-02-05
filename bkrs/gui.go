package bkrs

/*
	#include "util.h"
*/
import "C"
import "unsafe"

func MoveAppWindow(x int32, y int32) {
	C.moveAppWindow(C.int(x), C.int(y))
}

func MakeWindowAlwaysOnTop() {
	C.makeWindowAlwaysOnTop()
}

func LoadIcon() {
	C.loadIcon()
}

func MinimizeOnShortcut(window unsafe.Pointer) {
	C.minimizeOnShortcut(C.HWND(window))
}
