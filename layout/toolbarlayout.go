package layout

import (
	"fyne.io/fyne/v2"
)

type ToolbarObject struct {
	object fyne.CanvasObject
	width  float32
}

type ToolbarLayout struct {
	leftObjects   []ToolbarObject
	centerObjects []ToolbarObject
	rightObjects  []ToolbarObject
	padding       float32
	height        float32
}

func NewToolbarLayout(leftObjects,
	centerObjets,
	rightObjects []ToolbarObject, height, padding float32) *fyne.Container {

	t := &ToolbarLayout{}
	t.padding = padding
	t.height = height
	t.leftObjects = leftObjects
	t.centerObjects = centerObjets
	t.rightObjects = rightObjects

	var objects []fyne.CanvasObject

	for _, object := range t.leftObjects {
		objects = append(objects, object.object)
	}

	for _, object := range t.centerObjects {
		objects = append(objects, object.object)
	}

	for _, object := range t.rightObjects {
		objects = append(objects, object.object)
	}

	c := &fyne.Container{Layout: t, Objects: objects}

	return c
}

func (t *ToolbarLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, t.height)

	paddingTotal := (float32(len(objects)) - 1) * t.padding

	for _, object := range t.leftObjects {
		minSize.Width += object.width
	}

	for _, object := range t.centerObjects {
		minSize.Width += object.width
	}

	for _, object := range t.rightObjects {
		minSize.Width += object.width
	}

	minSize.Width += paddingTotal

	return minSize
}

func (t *ToolbarLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	var posX float32
	var centerBlocX, centerBlocWidth float32
	var rightBlocX, rightBlocWidth float32

	centerBlocWidth = t.getBlocWidth(t.centerObjects)
	rightBlocWidth = t.getBlocWidth(t.rightObjects)

	centerBlocX = (containerSize.Width / 2) - (centerBlocWidth / 2)
	rightBlocX = containerSize.Width - rightBlocWidth

	// Left objects position
	posX = 0

	for _, object := range t.leftObjects {
		object.object.Move(fyne.NewPos(posX, 0))
		object.object.Resize(fyne.NewSize(object.width, t.height))
		posX += object.width + t.padding
	}

	// Center objects position
	posX = centerBlocX

	for _, object := range t.centerObjects {
		object.object.Move(fyne.NewPos(posX, 0))
		object.object.Resize(fyne.NewSize(object.width, t.height))
		posX += object.width + t.padding
	}

	// Right objects position
	posX = rightBlocX

	for _, object := range t.rightObjects {
		object.object.Move(fyne.NewPos(posX, 0))
		object.object.Resize(fyne.NewSize(object.width, t.height))
		posX += object.width + t.padding
	}

}

func (t *ToolbarLayout) getBlocWidth(objects []ToolbarObject) float32 {
	var blocWidth float32

	for _, object := range objects {
		blocWidth += object.width
	}

	blocWidth += float32(len(objects)-1) * t.padding

	return blocWidth
}
