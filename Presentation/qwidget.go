package main

import "github.com/therecipe/qt/widgets"

//START OMIT
type Widget struct {
	widgets.QWidget

	_ func() `constructor:"init"`
}

//END OMIT
