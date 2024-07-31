package komok

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

var (
	shortcutCtrlA         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyA, fyne.KeyModifierControl)
	shortcutCtrlB         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyB, fyne.KeyModifierControl)
	shortcutCtrlH         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyH, fyne.KeyModifierControl)
	shortcutCtrlO         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyO, fyne.KeyModifierControl)
	shortcutCtrlS         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyS, fyne.KeyModifierControl)
	shortcutCtrlN         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyN, fyne.KeyModifierControl)
	shortcutCtrlW         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyW, fyne.KeyModifierControl)
	shortcutCtrlI         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyI, fyne.KeyModifierControl)
	shortcutCtrlD         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyD, fyne.KeyModifierControl)
	shortcutCtrlT         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyT, fyne.KeyModifierControl)
	shortcutCtrlE         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyE, fyne.KeyModifierControl)
	shortcutCtrlP         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyP, fyne.KeyModifierControl)
	shortcutCtrlF         *desktop.CustomShortcut = createCustomShortcut(fyne.KeyF, fyne.KeyModifierControl)
	shortcutCtrlUp        *desktop.CustomShortcut = createCustomShortcut(fyne.KeyUp, fyne.KeyModifierControl)
	shortcutCtrlDown      *desktop.CustomShortcut = createCustomShortcut(fyne.KeyDown, fyne.KeyModifierControl)
	shortcutCtrlTab       *desktop.CustomShortcut = createCustomShortcut(fyne.KeyTab, fyne.KeyModifierControl)
	shortcutCtrlEnter     *desktop.CustomShortcut = createCustomShortcut(fyne.KeyEnter, fyne.KeyModifierControl)
	shortcutCtrlShiftTab  *desktop.CustomShortcut = createCustomShortcut(fyne.KeyTab, fyne.KeyModifierControl|fyne.KeyModifierShift)
	shortcutCtrlShiftN    *desktop.CustomShortcut = createCustomShortcut(fyne.KeyN, fyne.KeyModifierControl|fyne.KeyModifierShift)
	shortcutCtrlShiftI    *desktop.CustomShortcut = createCustomShortcut(fyne.KeyI, fyne.KeyModifierControl|fyne.KeyModifierShift)
	shortcutCtrlShiftW    *desktop.CustomShortcut = createCustomShortcut(fyne.KeyW, fyne.KeyModifierControl|fyne.KeyModifierShift)
	shortcutCtrlShiftO    *desktop.CustomShortcut = createCustomShortcut(fyne.KeyO, fyne.KeyModifierControl|fyne.KeyModifierShift)
	shortcutCtrlShiftE    *desktop.CustomShortcut = createCustomShortcut(fyne.KeyE, fyne.KeyModifierControl|fyne.KeyModifierShift)
	shortcutCtrlShiftT    *desktop.CustomShortcut = createCustomShortcut(fyne.KeyT, fyne.KeyModifierControl|fyne.KeyModifierShift)
	shortcutCtrlShiftF    *desktop.CustomShortcut = createCustomShortcut(fyne.KeyF, fyne.KeyModifierControl|fyne.KeyModifierShift)
	shortcutCtrlShiftR    *desktop.CustomShortcut = createCustomShortcut(fyne.KeyR, fyne.KeyModifierControl|fyne.KeyModifierShift)
	shortcutCtrlShiftV    *desktop.CustomShortcut = createCustomShortcut(fyne.KeyV, fyne.KeyModifierControl|fyne.KeyModifierShift)
	shortcutCtrlShiftUp   *desktop.CustomShortcut = createCustomShortcut(fyne.KeyUp, fyne.KeyModifierControl|fyne.KeyModifierShift)
	shortcutCtrlShiftDown *desktop.CustomShortcut = createCustomShortcut(fyne.KeyDown, fyne.KeyModifierControl|fyne.KeyModifierShift)
)

// createCustomShortcut creates and return a pointer to a new custom shortcut
func createCustomShortcut(key fyne.KeyName, mod fyne.KeyModifier) *desktop.CustomShortcut {
	return &desktop.CustomShortcut{
		KeyName:  key,
		Modifier: mod,
	}
}
