package manager

import "fyne.io/fyne/v2"

type Shortcut struct {
	shortcut fyne.Shortcut
	action   func()
}

type ShortcutsManager struct {
	registredShortcuts []Shortcut
}

func NewShortcutsManager() *ShortcutsManager {
	return &ShortcutsManager{}
}

func (sm *ShortcutsManager) RegisterShortcut(shortcut fyne.Shortcut, action func()) {
	s := Shortcut{
		shortcut: shortcut,
		action:   action,
	}

	sm.registredShortcuts = append(sm.registredShortcuts, s)
}

func (sm *ShortcutsManager) TriggerShortcut(shortcut fyne.Shortcut) {
	a := sm.findShortcutAction(shortcut)

	if a == nil {
		return
	}

	a()
}

func (sm *ShortcutsManager) findShortcutAction(shortcut fyne.Shortcut) func() {
	sn := shortcut.ShortcutName()

	for _, v := range sm.registredShortcuts {
		if v.shortcut.ShortcutName() == sn {
			return v.action
		}
	}

	return nil
}
