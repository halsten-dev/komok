package manager

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/halsten-dev/komok/layout"
)

type ContentCode string

type IContent interface {
	GetCode() ContentCode

	// GetGUI returns the content to add to the window's content.
	GetGUI() fyne.CanvasObject

	// InitGUI is called after the GUI was added to the window's content.
	InitGUI()

	// Init is used to initialize all GUI elements.
	Init()

	// Destroy is used to destroy the GUI elements.
	Destroy()
}

type ContentManager struct {
	history  []ContentCode
	Contents []IContent

	CurrentContent     IContent
	CurrentContentCode ContentCode

	navbar      *fyne.Container
	navbarWidth float32

	app    fyne.App
	window fyne.Window
}

func NewContentManager(app fyne.App, window fyne.Window) *ContentManager {
	manager := &ContentManager{
		history:  make([]ContentCode, 0),
		Contents: make([]IContent, 0),
		app:      app,
		window:   window,
		navbar:   nil,
	}

	manager.window.SetContent(container.NewWithoutLayout())

	return manager
}

func (cm *ContentManager) SetNavbar(navbar *fyne.Container, width float32) {
	cm.navbar = navbar
	cm.navbarWidth = width
}

func (cm *ContentManager) RegisterContent(content IContent) {
	cm.Contents = append(cm.Contents, content)

	log.Printf("ContentManager : Content %s registred successfully", content.GetCode())
}

func (cm *ContentManager) ChangeContent(newCode ContentCode) {
	newContent := cm.GetContent(newCode)

	if newContent == nil {
		log.Printf("ContentManager : Content %s not found", newCode)
		return
	}

	if cm.CurrentContent != nil {
		cm.CurrentContent.Destroy()
		cm.addCurrentContentToHistory()
	}

	cm.window.SetContent(container.NewWithoutLayout())
	cm.CurrentContent = newContent
	cm.CurrentContentCode = newCode

	cm.CurrentContent.Init()

	if cm.navbar == nil {
		cm.window.SetContent(cm.CurrentContent.GetGUI())
	} else {
		navbar := container.New(layout.NewFixedAutoSizeLayout(cm.navbarWidth, true), cm.navbar,
			cm.CurrentContent.GetGUI())
		cm.window.SetContent(navbar)
	}

	cm.CurrentContent.InitGUI()
}

func (cm *ContentManager) addCurrentContentToHistory() {
	cm.history = append(cm.history, cm.CurrentContent.GetCode())

	log.Printf("ContentManager : Current content %s added to history", cm.CurrentContent.GetCode())
}

func (cm *ContentManager) GetContent(code ContentCode) IContent {
	for i, content := range cm.Contents {
		if content.GetCode() == code {
			return cm.Contents[i]
		}
	}

	return nil
}
