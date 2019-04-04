package main

import (
	"os"

	"github.com/therecipe/qt/gui"
)

func sss() {
	//START OMIT
	gui.NewQGuiApplication(len(os.Args), os.Args)
	//END OMIT
}
