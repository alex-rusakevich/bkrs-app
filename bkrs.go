package main

import (
	"bkrs/bkrs"
	"github.com/lxn/win"
	"github.com/webview/webview"
)

const (
	defaultW = 430
	defaultH = 550
)

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	bkrs.LoadConfig()
	w.Bind("isFeatureEnabled", bkrs.IsFeatureEnabled)
	w.Bind("getKeysByFeatureName", bkrs.GetKeysByFeatureName)
	w.Bind("setKeys", bkrs.SetKeys)
	w.Bind("setFeatureState", bkrs.SetFeatureState)
	defer bkrs.SaveConfig()

	settingsEnabled, _ := bkrs.IsFeatureEnabled("settings")
	if settingsEnabled {
		go bkrs.RunSettingsServer()
	}

	bkrs.MoveAppWindow(win.GetSystemMetrics(win.SM_CXSCREEN)-(defaultW+defaultW/2),
		win.GetSystemMetrics(win.SM_CYSCREEN)/2-defaultH/2)

	bkrs.MakeWindowAlwaysOnTop()
	go bkrs.MinimizeOnShortcut(w.Window())

	w.Init(bkrs.ReadFileIntoString("./resources/js/ui.js"))

	w.SetTitle("bkrs-app")
	w.SetSize(defaultW, defaultH, webview.HintNone)
	w.Navigate("https://bkrs.info")

	bkrs.LoadIcon()

	w.Run()
}
