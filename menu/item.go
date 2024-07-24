package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

type Item struct {
	ID                  string
	Label               string
	Shortcut            desktop.CustomShortcut
	Action              func()
	ActivationCondition func() bool
	Instance            *fyne.MenuItem
}

func newItem(id, label string, shortcut desktop.CustomShortcut,
	action func(), activationCondition func() bool) *Item {
	return &Item{
		ID:                  id,
		Label:               label,
		Shortcut:            shortcut,
		Action:              action,
		ActivationCondition: activationCondition,
	}
}
