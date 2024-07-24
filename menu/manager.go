package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/chicogamedev/komok"
)

type Manager struct {
	window fyne.Window
	scm    *komok.ShortcutsManager
	Menus  map[string]*Menu
}

func NewManager(win fyne.Window, scm *komok.ShortcutsManager) *Manager {
	return &Manager{
		window: win,
		scm:    scm,
		Menus:  map[string]*Menu{},
	}
}

func (m *Manager) CreateMenu(id, label string) {
	menu := newMenu(id, label)

	m.Menus[id] = menu
}

func (m *Manager) CreateMenuItem(menuID, id, label string, shortcut desktop.CustomShortcut, action func(),
	activationCondition func() bool) {
	item := newItem(id, label, shortcut, action, activationCondition)

	m.Menus[menuID].Items[id] = item
}

func (m *Manager) ConstructMainMenu() *fyne.MainMenu {
	var menus []*fyne.Menu

	for _, menu := range m.Menus {
		var items []*fyne.MenuItem

		for _, item := range menu.Items {
			item.Instance = fyne.NewMenuItem(item.Label, item.Action)
			item.Instance.Shortcut = &item.Shortcut
			m.scm.RegisterShortcut(&item.Shortcut, item.Action)
			m.window.Canvas().AddShortcut(&item.Shortcut, m.triggerCanvasShortcut)
			items = append(items, item.Instance)
		}

		menu.Instance = fyne.NewMenu(menu.Label, items...)
		menus = append(menus, menu.Instance)
	}

	return fyne.NewMainMenu(menus...)
}

func (m *Manager) UpdateMenu() {
	for _, menu := range m.Menus {
		for _, item := range menu.Items {
			active := item.ActivationCondition()
			item.Instance.Disabled = !active
		}
	}
}

func (m *Manager) triggerCanvasShortcut(shortcut fyne.Shortcut) {
	m.scm.TriggerShortcut(shortcut)
}
