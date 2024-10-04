package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

type orderCounter struct {
	// menuOrderCounter contains the current order for the main Menu bar and menuItemID
	// "main" key is reserved for the main Menu bar Menu order counter
	menuOrderCounter map[string]int

	// menuItemsOrderCounter contains the current order for any given menuID
	menuItemsOrderCounter map[string]int
}

func newOrderCounter() *orderCounter {
	oc := &orderCounter{
		menuOrderCounter:      make(map[string]int),
		menuItemsOrderCounter: make(map[string]int),
	}

	oc.menuOrderCounter["main"] = 0

	return oc
}

func (oc *orderCounter) getNextMenuOrder(itemID string) int {
	if itemID == "" {
		itemID = "main"
	}

	order := oc.menuOrderCounter[itemID]

	oc.menuOrderCounter[itemID]++

	return order
}

func (oc *orderCounter) getNextMenuItemOrder(menuID string) int {
	order := oc.menuItemsOrderCounter[menuID]

	oc.menuItemsOrderCounter[menuID]++

	return order
}

// Menu defines an individual Menu
type Menu struct {
	ID       string
	Label    string
	Order    int
	Instance *fyne.Menu
	IsChild  bool
	ItemID   string
}

// newMenu creates a *Menu and returns it
func newMenu(id, label string, order int) *Menu {
	return &Menu{
		ID:      id,
		Label:   label,
		Order:   order,
		IsChild: false,
	}
}

func newChildMenu(id, itemID string, order int) *Menu {
	return &Menu{
		ID:      id,
		Order:   order,
		IsChild: true,
		ItemID:  itemID,
	}
}

// MenuItem defines an individual Menu item
type MenuItem struct {
	ID                  string
	MenuID              string
	Label               string
	Order               int
	Shortcut            *desktop.CustomShortcut
	Action              func()
	ActivationCondition func() bool
	Instance            *fyne.MenuItem
	IsSeparator         bool
}

// newMenuItem creates a new Menu item and returns it as *MenuItem
func newMenuItem(id, menuID, label string, shortcut *desktop.CustomShortcut,
	action func(), activationCondition func() bool, order int) *MenuItem {
	return &MenuItem{
		ID:                  id,
		MenuID:              menuID,
		Label:               label,
		Order:               order,
		Shortcut:            shortcut,
		Action:              action,
		ActivationCondition: activationCondition,
		IsSeparator:         false,
	}
}

func newMenuItemSeparator(menuID string, order int) *MenuItem {
	return &MenuItem{
		MenuID:      menuID,
		Order:       order,
		IsSeparator: true,
	}
}
