package menu

import "fyne.io/fyne/v2"

type Menu struct {
	ID       string
	Label    string
	Items    map[string]*Item
	Instance *fyne.Menu
}

func newMenu(id, label string) *Menu {
	return &Menu{
		ID:    id,
		Label: label,
		Items: make(map[string]*Item),
	}
}
