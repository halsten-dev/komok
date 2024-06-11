package komok

import (
	"fyne.io/fyne/v2"
)

type AutoSize struct {
	// If it's a horizontal autosize.
	horizontal bool
}

func NewAutoSizeLayout(isHorizontal bool) fyne.Layout {
	return &AutoSize{isHorizontal}
}

func (as *AutoSize) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		minSize = minSize.Max(child.MinSize())
	}

	return minSize
}

func (as *AutoSize) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	topLeft := fyne.NewPos(0, 0)

	var newSize fyne.Size
	var currentSize float32

	for _, child := range objects {
		if as.horizontal {
			currentSize = child.Size().Height

			if currentSize <= 0 {
				currentSize = containerSize.Height
			}

			newSize = fyne.NewSize(containerSize.Width, currentSize)
		} else {
			currentSize = child.Size().Width

			if currentSize <= 0 {
				currentSize = containerSize.Width
			}

			newSize = fyne.NewSize(currentSize, containerSize.Height)
		}

		child.Resize(newSize)
		child.Move(topLeft)
	}
}
