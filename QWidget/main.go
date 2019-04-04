package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var (
		window = widgets.NewQMainWindow(nil, 0)
		vbox   = widgets.NewQVBoxLayout2(nil)
		button = widgets.NewQPushButton2("click me!", nil)
	)
	vbox.AddWidget(button, 0, 0)

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(vbox)
	window.SetCentralWidget(centralWidget)

	window.Show()
	widgets.QApplication_Exec()
}
