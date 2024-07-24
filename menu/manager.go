package menu

import (
	"fyne.io/fyne/v2"
	"github.com/chicogamedev/komok"
)

type manager struct {
	window *fyne.Window
	scm    *komok.ShortcutsManager
	Menus  map[string]*Menu
}

func NewManager(win *fyne.Window, scm *komok.ShortcutsManager) *manager {
	return &manager{
		window: win,
		scm:    scm,
	}
}

func (m *manager) CreateMenu(id, label string) {
	menu := newMenu(id, label)

	m.Menus[id] = menu
}

func (m *manager) CreateMenuItem(menuID, id, label string, shortcut fyne.Shortcut, action func(),
	activationCondition func() bool) {
	item := newItem(id, label, shortcut, action, activationCondition)

	m.Menus[menuID].Items[id] = item
}

func (m *manager) ConstructMainMenu() *fyne.MainMenu {
	var menus []*fyne.Menu

	for _, menu := range m.Menus {
		var items []*fyne.MenuItem

		for _, item := range menu.Items {
			item.Instance = fyne.NewMenuItem(item.Label, item.Action)
			m.scm.RegisterShortcut(item.Shortcut, item.Action)
			m.window.Canvas().AddShortcut(item.Shortcut, m.triggerCanvasShortcut)
			items = append(items, item.Instance)
		}

		menu.Instance = fyne.NewMenu(menu.Label, items...)
		menus = append(menus, menu.Instance)
	}

	return fyne.NewMainMenu(menus...)
}

func (m *manager) UpdateMenu() {
	for _, menu := range m.Menus {
		for _, item := range menu.Items {
			active := item.ActivationCondition()
			item.Instance.Disabled = !active
		}
	}
}

func (m *manager) triggerCanvasShortcut(shortcut fyne.Shortcut) {
	m.scm.TriggerShortcut(shortcut)
}
