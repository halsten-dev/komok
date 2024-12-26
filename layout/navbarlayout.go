package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"golang.org/x/image/colornames"
)

type NavbarLayout struct {
	topObjects    []fyne.CanvasObject
	centerObjects []fyne.CanvasObject
	bottomObjects []fyne.CanvasObject
	padding       float32
	horizontal    bool
	separator     *canvas.Line
}

func NewNavbarLayout(
	topObjects,
	centerObjets,
	bottomObjects []fyne.CanvasObject, padding float32) *fyne.Container {

	t := &NavbarLayout{}
	t.padding = padding
	t.topObjects = topObjects
	t.centerObjects = centerObjets
	t.bottomObjects = bottomObjects

	var objects []fyne.CanvasObject

	for _, object := range t.topObjects {
		objects = append(objects, object)
	}

	for _, object := range t.centerObjects {
		objects = append(objects, object)
	}

	for _, object := range t.bottomObjects {
		objects = append(objects, object)
	}

	separator := canvas.NewLine(colornames.Black)

	objects = append(objects, separator)

	c := &fyne.Container{Layout: t, Objects: objects}

	return c
}

func (t *NavbarLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)

	paddingTotal := (float32(len(objects)) - 1) * t.padding

	for _, object := range objects {
		minSize.Width += object.MinSize().Width
		minSize.Height += object.MinSize().Height
	}

	minSize.Width += paddingTotal

	return minSize
}

func (t *NavbarLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	var posY float32
	var centerBlocY, centerBlocHeight float32
	var bottomBlocY, bottomBlocHeight float32
	var objectWidth float32

	objectWidth = containerSize.Width - 5

	centerBlocHeight = t.getBlocHeight(t.centerObjects)
	bottomBlocHeight = t.getBlocHeight(t.bottomObjects)

	centerBlocY = (containerSize.Height / 2) - (centerBlocHeight / 2)
	bottomBlocY = containerSize.Height - bottomBlocHeight

	// Left objects position
	posY = 0

	for _, object := range t.topObjects {
		object.Move(fyne.NewPos(0, posY))
		object.Resize(fyne.NewSize(objectWidth, object.MinSize().Height))
		posY += object.MinSize().Height + t.padding
	}

	// Center objects position
	posY = centerBlocY

	for _, object := range t.centerObjects {
		object.Move(fyne.NewPos(0, posY))
		object.Resize(fyne.NewSize(objectWidth, object.MinSize().Height))
		posY += object.MinSize().Height + t.padding
	}

	// Right objects position
	posY = bottomBlocY

	for _, object := range t.bottomObjects {
		object.Move(fyne.NewPos(0, posY))
		object.Resize(fyne.NewSize(objectWidth, object.MinSize().Height))
		posY += object.MinSize().Height + t.padding
	}

	t.separator.Position1 = fyne.NewPos(containerSize.Width-1, 0)
	t.separator.Position2 = fyne.NewPos(containerSize.Width, containerSize.Height)
}

func (t *NavbarLayout) getBlocHeight(objects []fyne.CanvasObject) float32 {
	var blocHeight float32

	for _, object := range objects {
		blocHeight += object.MinSize().Height
	}

	blocHeight += float32(len(objects)-1) * t.padding

	return blocHeight
}
