package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type CombatEncounterTab struct {
	widgets.QWidget

	_ func() `constructor:"init"`

	characterList       core.QList
	characterListWidget *widgets.QListWidget

	layout *widgets.QVBoxLayout
}

func (t *CombatEncounterTab) init() {
	t.layout = widgets.NewQVBoxLayout()

	// Character list widget
	t.characterListWidget = widgets.NewQListWidget(nil)
	t.characterListWidget.AddItem("Param")

	// Character list
	//core.QList

	// Character layout
	characterLayout := widgets.NewQVBoxLayout()
	characterLayout.AddWidget(t.characterListWidget, 0, 0)

	///////////////////////////////////////////////////////////////////////////

	//characterListWidget.Connect

	// Character view
	//characterView := widgets.Model

	// Character layout
	//characterLayout := widgets.NewQVBoxLayout()
	//characterLayout.AddWidget(characterListWidget, 0, 0)

	// Character button
	addCharacterButton := widgets.NewQPushButton2("Add Character", nil)
	addCharacterButton.ConnectClicked(func(checked bool) {
		t.characterListWidget.AddItem("clicked heee")
	})

	// Dice button
	rollDiceButton := widgets.NewQPushButton2("Roll Dice", nil)
	rollDiceButton.ConnectClicked(func(checked bool) {
		widgets.QMessageBox_Information(nil, "Warning", "Not implemented functionality",
			widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	// Button layout
	buttonLayout := widgets.NewQHBoxLayout()
	buttonLayout.AddWidget(addCharacterButton, 0, 0)
	buttonLayout.AddWidget(rollDiceButton, 0, 0)

	// Add layouts to CombatEncounterTab
	t.layout.AddLayout(characterLayout, 0)
	t.layout.AddLayout(buttonLayout, 0)

	t.SetLayout(t.layout)

}
