package komok

import (
	"fyne.io/fyne/v2"
	"github.com/halsten-dev/komok/manager"
	"github.com/halsten-dev/komok/manager/menu"
)

var AppEngine *Engine

type Engine struct {
	Widgets         map[string]fyne.Widget
	Systems         map[string]System
	MenuManager     *menu.Manager
	ShortcutManager *manager.ShortcutsManager
	Window          fyne.Window
}

type System interface {
	HandleSignal(e *Engine, s string)
}

func InitAppEngine(window fyne.Window) {
	sm := manager.NewShortcutsManager()
	AppEngine = &Engine{
		Widgets:         make(map[string]fyne.Widget),
		Systems:         make(map[string]System),
		ShortcutManager: sm,
		MenuManager:     menu.NewManager(window, sm),
		Window:          window,
	}
}

func (k *Engine) RegisterWidget(code string, widget fyne.Widget) {
	k.Widgets[code] = widget
}

func (k *Engine) RegisterSystem(code string, system System) {
	k.Systems[code] = system
}

func (k *Engine) SendSignal(s string) {
	for id := range k.Systems {
		k.Systems[id].HandleSignal(k, s)
	}
}
