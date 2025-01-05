package komok

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

var (
	ShortcutCtrlA      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyA, fyne.KeyModifierControl)
	ShortcutCtrlB      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyB, fyne.KeyModifierControl)
	ShortcutCtrlH      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyH, fyne.KeyModifierControl)
	ShortcutCtrlO      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyO, fyne.KeyModifierControl)
	ShortcutCtrlS      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyS, fyne.KeyModifierControl)
	ShortcutCtrlN      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyN, fyne.KeyModifierControl)
	ShortcutCtrlW      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyW, fyne.KeyModifierControl)
	ShortcutCtrlI      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyI, fyne.KeyModifierControl)
	ShortcutCtrlD      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyD, fyne.KeyModifierControl)
	ShortcutCtrlT      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyT, fyne.KeyModifierControl)
	ShortcutCtrlE      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyE, fyne.KeyModifierControl)
	ShortcutCtrlP      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyP, fyne.KeyModifierControl)
	ShortcutCtrlF      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyF, fyne.KeyModifierControl)
	ShortcutCtrlUp     *desktop.CustomShortcut = createCustomShortcut(fyne.KeyUp, fyne.KeyModifierControl)
	ShortcutCtrlDown   *desktop.CustomShortcut = createCustomShortcut(fyne.KeyDown, fyne.KeyModifierControl)
	ShortcutCtrlTab    *desktop.CustomShortcut = createCustomShortcut(fyne.KeyTab, fyne.KeyModifierControl)
	ShortcutCtrlEnter  *desktop.CustomShortcut = createCustomShortcut(fyne.KeyEnter, fyne.KeyModifierControl)
	ShortcutCtrlReturn *desktop.CustomShortcut = createCustomShortcut(fyne.KeyReturn, fyne.KeyModifierControl)

	ShortcutCtrlShiftTab *desktop.CustomShortcut = createCustomShortcut(fyne.KeyTab,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftN *desktop.CustomShortcut = createCustomShortcut(fyne.KeyN,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftI *desktop.CustomShortcut = createCustomShortcut(fyne.KeyI,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftW *desktop.CustomShortcut = createCustomShortcut(fyne.KeyW,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftO *desktop.CustomShortcut = createCustomShortcut(fyne.KeyO,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftE *desktop.CustomShortcut = createCustomShortcut(fyne.KeyE,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftT *desktop.CustomShortcut = createCustomShortcut(fyne.KeyT,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftF *desktop.CustomShortcut = createCustomShortcut(fyne.KeyF,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftR *desktop.CustomShortcut = createCustomShortcut(fyne.KeyR,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftV *desktop.CustomShortcut = createCustomShortcut(fyne.KeyV,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftUp *desktop.CustomShortcut = createCustomShortcut(fyne.KeyUp,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftDown *desktop.CustomShortcut = createCustomShortcut(fyne.KeyDown,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
	ShortcutCtrlShiftReturn *desktop.CustomShortcut = createCustomShortcut(fyne.KeyReturn,
		fyne.KeyModifierControl|fyne.KeyModifierShift)
)

// createCustomShortcut creates and return a pointer to a new custom shortcut
func createCustomShortcut(key fyne.KeyName, mod fyne.KeyModifier) *desktop.CustomShortcut {
	var shortcut *desktop.CustomShortcut

	shortcut = &desktop.CustomShortcut{
		KeyName:  key,
		Modifier: mod,
	}

	return shortcut
}
