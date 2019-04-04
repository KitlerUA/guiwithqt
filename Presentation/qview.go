package main

func view() {
	//START OMIT
	var view = quick.NewQQuickView(nil)

	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))

	view.Show()
	//END OMIT
}