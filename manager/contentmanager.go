package manager

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"log"
)

type ContentCode string

type IContent interface {
	GetCode() ContentCode
	GetGUI() fyne.CanvasObject
	Init()
	Destroy()
}

type ContentManager struct {
	history  []ContentCode
	Contents []IContent

	CurrentContent IContent

	app    fyne.App
	window fyne.Window
}

var Manager *ContentManager

func Init(app fyne.App, window fyne.Window) {
	Manager = &ContentManager{
		history:  make([]ContentCode, 0),
		Contents: make([]IContent, 0),
		app:      app,
		window:   window,
	}

	Manager.window.SetContent(container.NewWithoutLayout())
}

func (cm *ContentManager) RegisterContent(content IContent) {
	cm.Contents = append(cm.Contents, content)

	log.Printf("ContentManager : Content %s registred successfully", content.GetCode())
}

func (cm *ContentManager) ChangeContent(newCode ContentCode) {
	newContent := cm.getContent(newCode)

	if newContent == nil {
		log.Printf("ContentManager : Content %s not found", newCode)
		return
	}

	cm.CurrentContent.Destroy()
	cm.addCurrentContentToHistory()

	cm.window.SetContent(container.NewWithoutLayout())
	cm.CurrentContent = newContent

	cm.CurrentContent.Init()

	cm.window.SetContent(container.NewStack(cm.CurrentContent.GetGUI()))
}

func (cm *ContentManager) addCurrentContentToHistory() {
	cm.history = append(cm.history, cm.CurrentContent.GetCode())

	log.Printf("ContentManager : Current content %s added to history", cm.CurrentContent.GetCode())
}

func (cm *ContentManager) getContent(code ContentCode) IContent {
	for i, content := range cm.Contents {
		if content.GetCode() == code {
			return cm.Contents[i]
		}
	}

	return nil
}
