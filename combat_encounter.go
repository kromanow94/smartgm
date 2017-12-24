package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type CombatEncounterTab struct {
	widgets.QWidget
	characterList       core.QList
	characterListWidget *CharacterWidget
	//characterListWidget *widgets.QListWidget
	layout              *widgets.QVBoxLayout

	_ func() `constructor:"init"`
}

func (t *CombatEncounterTab) init() {
	t.layout = widgets.NewQVBoxLayout()


	// Character list widget
	t.characterListWidget = NewCharacterWidget(nil, 0)
	//t.characterListWidget = widgets.NewQListWidget(nil)
	//t.characterListWidget.AddItem("Param")
	asd := make([][]string, 2)

	asd[0] = make([]string, 2)
	asd[1] = make([]string, 2)

	asd[0][0] = "tuti 00"
	asd[0][1] = "tuti 01"
	asd[1][0] = "tuti 10"
	asd[1][1] = "tuti 11"

	t.characterListWidget.AddCharacter("Och", asd)


	// Character layout
	//characterLayout := widgets.NewQVBoxLayout()
	//characterLayout.AddWidget(t.characterListWidget, 0, 0)

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
		//t.characterListWidget.AddItem("clicked heee")
		t.characterListWidget.AddCharacter("Clicked", asd)
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
	t.layout.AddWidget(t.characterListWidget, 0, 0)
	t.layout.AddLayout(buttonLayout, 0)

	t.SetLayout(t.layout)

}
