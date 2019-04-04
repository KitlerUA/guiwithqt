package main

import (
	"github.com/therecipe/qt/core"
)

//START OMIT
type Person struct {
	core.QObject

	_ string `property:"name"`
	_ int    `property:"shoeSize"`
}

//END OMIT
