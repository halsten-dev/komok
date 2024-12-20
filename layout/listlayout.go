package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

type ListLayout struct {
	visibleItemCount int
}

func NewListLayout(visibleItemCount int) *ListLayout {
	return &ListLayout{visibleItemCount: visibleItemCount}
}

func (l *ListLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)

	list, ok := objects[0].(*widget.List)

	if !ok {
		log.Println("Komok : ListLayout should only hold a list")
		return minSize
	}

	// List minSize is the height of 1 item
	itemHeight := list.MinSize().Height

	minSize.Height = itemHeight*float32(l.visibleItemCount) + (theme.Padding() * float32(l.visibleItemCount))
	minSize.Width = list.MinSize().Width

	return minSize
}

func (l *ListLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	list, ok := objects[0].(*widget.List)

	if !ok {
		log.Println("Komok : ListLayout should only hold a list")
		return
	}

	list.Move(fyne.NewPos(0, 0))
	list.Resize(containerSize)
}
