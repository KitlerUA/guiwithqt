package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

func main() {
	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetApplicationDisplayName("Hallo freunde!")
	fmt.Println(app.ApplicationDirPath())
	//app.Quit()
	window := NewRasterWindow(nil)
	window.Show()
	app.Exec()
}

type RasterWindow struct {
	gui.QWindow

	_ func() `constructor:"init"`

	render func(painter *gui.QPainter)

	m_backingStore *gui.QBackingStore
}

func (w *RasterWindow) init() {
	w.Create()
	w.m_backingStore = gui.NewQBackingStore(w)

	w.SetGeometry(100, 100, 300, 200)

	w.render = w.renderFunc
}

func (w *RasterWindow) renderFunc(painter *gui.QPainter) {
	painter.DrawText5(core.NewQRectF4(0, 0, float64(w.Width()), float64(w.Height())), int(core.Qt__AlignCenter), "QWindow", core.NewQRectF())
}
