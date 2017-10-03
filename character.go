package main

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"go/ast"
)

type Character struct {
	name    string
	details DetailsTableView
}

type DetailsTableView struct {
	widgets.QTableView
	rows [][]string
}

type CharacterWidget struct {
	//core.QAbstractListModel
	widgets.QWidget

	_ func() `constructor:"init"`

	// Layout of this widget
	layout *widgets.QHBoxLayout

	// first pane - character names widget
	characterListWidget *widgets.QListWidget

	// second pane - details
	detailsWidget *widgets.QTableView

	// third pane - options
	//??

	_ map[string]*Character `property:"objects"`
	_ []*Character `property:"characters"`
	characters []*Character

}

func (cw *CharacterWidget) init() {
	cw.layout = widgets.NewQHBoxLayout()

	cw.characterListWidget = widgets.NewQListWidget(nil)
	cw.detailsWidget = widgets.NewQTableView(nil)



	// character for tests
	//details := DetailsTableView{[][]string{
	//	{"Health", "20"},
	//	{"k20", "2"},
	//}}
	//character := Character{"Param", details}
	//characters := []*Character{}

	// tests
	cw.ConnectCharacters(func() []*Character {
		return cw.characters
	})


	cw.SetLayout(cw.layout)
}

func (cw *CharacterWidget) AddCharacter(name string, stringRows [][]string) {
	character := &Character{}
	character.name = name

	details := DetailsTableView{}
	details.rows = stringRows

	tableModel := core.NewQAbstractTableModel(nil)
	tableModel.ConnectRowCount(func(parent *core.QModelIndex) int {
		return len(details.rows)
	})
	tableModel.ConnectColumnCount(func(parent *core.QModelIndex) int {
		return 2
	})
	tableModel.ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		if role != int(core.Qt__DisplayRole) {
			return core.NewQVariant()
		}
		return core.NewQVariant14(details.rows[index.Row()][index.Column()])
	})

	details.SetModel(tableModel)

	character.details = details

	cw.characters = append(cw.characters, character)


	//cw.characterListWidget.AddItem(character.name)
	//cw.detailsWidget.
}