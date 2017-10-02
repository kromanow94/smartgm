package main

import (
    "os"

    "github.com/therecipe/qt/widgets"
 )

func main() {

    // Create application
    app := widgets.NewQApplication(len(os.Args), os.Args)

    // Create main window
    window := widgets.NewQMainWindow(nil, 0)
    window.SetWindowTitle("Smart Game Master")
    window.SetMinimumSize2(800, 600)

    // Create main window layout
    layout := widgets.NewQVBoxLayout()

    // Create main widget and set the layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(layout)

	// Create tab widget
	tabWidget := widgets.NewQTabWidget(nil)
	tabWidget.SetWindowTitle("Various tabs")
	//tabWidget.SetSizePolicy(widgets.NewQSizePolicy2())

	// Create CombatEncounter tab
	combatEncounterTab := NewCombatEncounterTab(nil, 0)
	tabWidget.AddTab(combatEncounterTab, "Combat Encounter")

	// Create just some list tab
	someList := widgets.NewQListWidget(nil)
	someList.AddItem("Hehee")
	tabWidget.AddTab(someList, "Some list")

	layout.AddWidget(tabWidget, 0, 0)


	//// Create a line edit and add it to the layout
	//input := widgets.NewQLineEdit(nil)
	//input.SetPlaceholderText("1. Write Something")
	//layout.AddWidget(input,0, 0)
	//
	//// Create a button and add it to the layout
	//button := widgets.NewQPushButton2("2. Click Me", nil)
	//layout.AddWidget(button, 0, 0)
	//
	//button.ConnectClicked(func(checked bool) {
	//	widgets.QMessageBox_Information(nil, "OK", input.Text(),
	//		widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	//})

	// Set main widget as the central widget of the window
	window.SetCentralWidget(mainWidget)


	// Show the window
	window.Show()


	// Execute app
    app.Exec()

}
