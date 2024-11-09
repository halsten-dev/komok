package menu

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/halsten-dev/komok/manager"
	"sort"
)

// Manager is the holder of everything concerning the applications menus
type Manager struct {
	window       fyne.Window
	scm          *manager.ShortcutsManager
	Menus        map[string]*Menu
	MenusItems   map[string]*MenuItem
	OrderCounter *orderCounter
}

// NewManager takes the window and the shortcutsManager and initialize a new *Manager
func NewManager(win fyne.Window, scm *manager.ShortcutsManager) *Manager {
	return &Manager{
		window:       win,
		scm:          scm,
		Menus:        map[string]*Menu{},
		MenusItems:   map[string]*MenuItem{},
		OrderCounter: newOrderCounter(),
	}
}

// CreateMenu creates a new menu with the given id and label
func (m *Manager) CreateMenu(id, label string) {
	menu := newMenu(id, label, m.OrderCounter.getNextMenuOrder(""))

	m.Menus[id] = menu

	// Init items counter
	m.OrderCounter.menuItemsOrderCounter[id] = 0
}

// CreateChildMenu creates a new menu with the given id and label
func (m *Manager) CreateChildMenu(itemID, id string) {
	menu := newChildMenu(id, itemID, m.OrderCounter.getNextMenuOrder(itemID))

	m.Menus[id] = menu

	// Init items counter
	m.OrderCounter.menuItemsOrderCounter[id] = 0
}

// CreateMenuItem creates a new menu item for the given menu
func (m *Manager) CreateMenuItem(menuID, id, label string, shortcut *desktop.CustomShortcut, action func(),
	activationCondition func() bool) {

	item := newMenuItem(id, menuID, label, shortcut, action, activationCondition,
		m.OrderCounter.getNextMenuItemOrder(menuID))
	m.MenusItems[id] = item

	// Init menu counter with the itemID (for possible child menu)
	m.OrderCounter.menuOrderCounter[id] = 0
}

// CreateMenuSeparator creates a new menu item separator
func (m *Manager) CreateMenuSeparator(menuID string) {
	order := m.OrderCounter.getNextMenuItemOrder(menuID)
	item := newMenuItemSeparator(menuID, order)
	m.MenusItems[fmt.Sprintf("%s#sep#%v", menuID, order)] = item
}

// GetMenu returns the menu for the given menuID
func (m *Manager) GetMenu(menuID string) *Menu {
	return m.Menus[menuID]
}

func (m *Manager) GetMenuItem(itemID string) *MenuItem {
	return m.MenusItems[itemID]
}

// getMenus returns the list of menus that are not children
func (m *Manager) getMenus() []*Menu {
	var menus []*Menu

	for _, menu := range m.Menus {
		if !menu.IsChild {
			menus = append(menus, menu)
		}
	}

	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Order < menus[j].Order
	})

	return menus
}

// getChildMenus returns the list of menus tagged as children
func (m *Manager) getChildMenus() []*Menu {
	var menus []*Menu

	for _, menu := range m.Menus {
		if menu.IsChild {
			menus = append(menus, menu)
		}
	}

	return menus
}

// getMenuItems returns the list of items for the given menuID
func (m *Manager) getMenuItems(menuID string) []*MenuItem {
	var items []*MenuItem

	for _, item := range m.MenusItems {
		if item.MenuID == menuID {
			items = append(items, item)
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Order < items[j].Order
	})

	return items
}

// ConstructMainMenu is used to build the fyne.MainMenu based on the Manager definition
func (m *Manager) ConstructMainMenu() *fyne.MainMenu {
	var menus []*fyne.Menu

	for _, menu := range m.getMenus() {
		var items []*fyne.MenuItem

		for _, item := range m.getMenuItems(menu.ID) {
			if item.IsSeparator {
				items = append(items, fyne.NewMenuItemSeparator())
				continue
			}

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

	childMenus := m.getChildMenus()

	if len(childMenus) > 0 {
		for _, childMenu := range childMenus {
			var childMenuItems []*fyne.MenuItem

			for _, childItem := range m.getMenuItems(childMenu.ID) {
				if childItem.IsSeparator {
					childMenuItems = append(childMenuItems, fyne.NewMenuItemSeparator())
					continue
				}

				childItem.Instance = fyne.NewMenuItem(childItem.Label, childItem.Action)

				if childItem.Shortcut != nil {
					childItem.Instance.Shortcut = childItem.Shortcut
					m.scm.RegisterShortcut(childItem.Shortcut, childItem.Action)
					m.window.Canvas().AddShortcut(childItem.Shortcut, m.triggerCanvasShortcut)
				}

				childMenuItems = append(childMenuItems, childItem.Instance)
			}

			menuInstance := fyne.NewMenu("", childMenuItems...)
			m.MenusItems[childMenu.ItemID].Instance.ChildMenu = menuInstance
			childMenu.Instance = menuInstance
		}
	}

	return fyne.NewMainMenu(menus...)
}

// UpdateMenu goes through all menus items to check if they need to be activated or not
func (m *Manager) UpdateMenu() {
	for _, item := range m.MenusItems {
		if item.ActivationCondition != nil {
			active := item.ActivationCondition()
			item.Instance.Disabled = !active
		}
	}
}

func (m *Manager) triggerCanvasShortcut(shortcut fyne.Shortcut) {
	m.scm.TriggerShortcut(shortcut)
}
