package manager

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

// menu defines an individual menu
type menu struct {
	ID           string
	Label        string
	Items        map[string]*menuItem
	ItemsIDOrder []string
	Instance     *fyne.Menu
}

// newMenu creates a *menu and returns it
func newMenu(id, label string) *menu {
	return &menu{
		ID:    id,
		Label: label,
		Items: make(map[string]*menuItem),
	}
}

// menuItem defines an individual menu item
type menuItem struct {
	ID                  string
	Label               string
	Shortcut            *desktop.CustomShortcut
	Action              func()
	ActivationCondition func() bool
	Instance            *fyne.MenuItem
}

// newMenuItem creates a new menu item and returns it as *menuItem
func newMenuItem(id, label string, shortcut *desktop.CustomShortcut,
	action func(), activationCondition func() bool) *menuItem {
	return &menuItem{
		ID:                  id,
		Label:               label,
		Shortcut:            shortcut,
		Action:              action,
		ActivationCondition: activationCondition,
	}
}

// MenuManager is the holder of everything concerning the applications menus
type MenuManager struct {
	window       fyne.Window
	scm          *ShortcutsManager
	Menus        map[string]*menu
	MenusIDOrder []string
}

// NewMenuManager takes the window and the shortcutsManager and initialize a new *MenuManager
func NewMenuManager(win fyne.Window, scm *ShortcutsManager) *MenuManager {
	return &MenuManager{
		window: win,
		scm:    scm,
		Menus:  map[string]*menu{},
	}
}

// CreateMenu creates a new menu with the given id and label
func (m *MenuManager) CreateMenu(id, label string) {
	menu := newMenu(id, label)

	m.Menus[id] = menu
	m.MenusIDOrder = append(m.MenusIDOrder, id)
}

// CreateMenuItem creates a new menu item for the given menu
func (m *MenuManager) CreateMenuItem(menuID, id, label string, shortcut *desktop.CustomShortcut, action func(),
	activationCondition func() bool) {

	item := newMenuItem(id, label, shortcut, action, activationCondition)
	m.Menus[menuID].Items[id] = item
	m.Menus[menuID].ItemsIDOrder = append(m.Menus[menuID].ItemsIDOrder, id)
}

func (m *MenuManager) CreateMenuSeparator(menuID string) {
	m.Menus[menuID].ItemsIDOrder = append(m.Menus[menuID].ItemsIDOrder, "separator")
}

// ConstructMainMenu is use to build the fyne.MainMenu based on the MenuManager definition
func (m *MenuManager) ConstructMainMenu() *fyne.MainMenu {
	var menus []*fyne.Menu

	for _, menuID := range m.MenusIDOrder {
		var items []*fyne.MenuItem
		menu := m.Menus[menuID]

		for _, itemID := range menu.ItemsIDOrder {
			if itemID == "separator" {
				items = append(items, fyne.NewMenuItemSeparator())
				continue
			}

			item := menu.Items[itemID]

			item.Instance = fyne.NewMenuItem(item.Label, item.Action)

			if item.Shortcut != nil {
				item.Instance.Shortcut = item.Shortcut
				m.scm.RegisterShortcut(item.Shortcut, item.Action)
				m.window.Canvas().AddShortcut(item.Shortcut, m.triggerCanvasShortcut)
			}

			items = append(items, item.Instance)
		}

		menu.Instance = fyne.NewMenu(menu.Label, items...)
		menus = append(menus, menu.Instance)
	}

	return fyne.NewMainMenu(menus...)
}

// UpdateMenu goes through all menus items to check if they need to be activated or not
func (m *MenuManager) UpdateMenu() {
	for _, menu := range m.Menus {
		for _, item := range menu.Items {
			active := item.ActivationCondition()
			item.Instance.Disabled = !active
		}
	}
}

func (m *MenuManager) triggerCanvasShortcut(shortcut fyne.Shortcut) {
	m.scm.TriggerShortcut(shortcut)
}
