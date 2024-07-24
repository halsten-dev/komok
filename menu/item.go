package menu

import "fyne.io/fyne/v2"

type Item struct {
	ID                  string
	Label               string
	Shortcut            fyne.Shortcut
	Action              func()
	ActivationCondition func() bool
	Instance            *fyne.MenuItem
}

func newItem(id, label string, shortcut fyne.Shortcut,
	action func(), activationCondition func() bool) *Item {
	return &Item{
		ID:                  id,
		Text:                label,
		Shortcut:            shortcut,
		Action:              action,
		ActivationCondition: activationCondition,
	}
}
