package layout

import (
	"fyne.io/fyne/v2"
)

type CenteredLayout struct {
	fixedWidth float32
}

func NewCenteredLayout(fixedWidth float32) fyne.Layout {
	return &CenteredLayout{
		fixedWidth: fixedWidth,
	}
}

func (c *CenteredLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	var height float32

	minSize := fyne.NewSize(c.fixedWidth, 0)

	for _, child := range objects {
		height += child.MinSize().Height
	}

	return minSize
}

func (c *CenteredLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	var pos fyne.Position
	var height float32

	topLeft := fyne.NewPos(
		containerSize.Width/2-c.fixedWidth/2,
		containerSize.Height/2-c.MinSize(objects).Height/2,
	)

	for i, child := range objects {
		if i == 0 {
			pos = topLeft
		} else {
			pos = fyne.NewPos(topLeft.X, topLeft.Y+height)
		}

		child.Move(pos)
		child.Resize(fyne.NewSize(c.fixedWidth, child.Size().Height))

		height += child.Size().Height
	}
}
