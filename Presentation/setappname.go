package main

import (
	"os"

	"github.com/therecipe/qt/gui"
)

func qwert() {
	//START OMIT
	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	app.SetApplicationDisplayName("Quatrix Express")

	gui.QGuiApplication_Exec()
	//END OMIT
}
