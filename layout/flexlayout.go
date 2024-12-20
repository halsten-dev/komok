package layout

import (
	"fyne.io/fyne/v2"
)

type FlexLayout struct {
	isHorizontal bool
	padding      float32
}

func NewFlexLayout(padding float32, isHorizontal bool) *FlexLayout {
	return &FlexLayout{
		padding:      padding,
		isHorizontal: isHorizontal,
	}
}

func (f *FlexLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)

	for _, object := range objects {
		if f.isHorizontal {
			if minSize.Height < object.MinSize().Height {
				minSize.Height = object.MinSize().Height
			}

			minSize.Width += object.MinSize().Width
		} else {
			if minSize.Width < object.MinSize().Width {
				minSize.Width = object.MinSize().Width
			}

			minSize.Height += object.MinSize().Height
		}
	}

	return minSize
}

func (f *FlexLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	var objectSize fyne.Size
	var usableSize fyne.Size

	if f.isHorizontal {
		usableSize = fyne.NewSize(
			containerSize.Width-(float32(len(objects)-1)*f.padding),
			containerSize.Height,
		)

		objectSize = fyne.NewSize(
			usableSize.Width/float32(len(objects)), usableSize.Height)
	} else {
		usableSize = fyne.NewSize(
			containerSize.Width,
			containerSize.Height-(float32(len(objects)-1)*f.padding),
		)

		objectSize = fyne.NewSize(
			usableSize.Width, usableSize.Height/float32(len(objects)),
		)
	}

	topLeft := fyne.NewPos(0, 0)

	for i, object := range objects {
		if f.isHorizontal {
			object.Resize(objectSize)
			object.Move(fyne.NewPos(topLeft.X+((objectSize.Width+f.padding)*float32(i)), topLeft.Y))
		} else {
			object.Resize(objectSize)
			object.Move(fyne.NewPos(topLeft.X, topLeft.Y+((objectSize.Height+f.padding)*float32(i))))
		}
	}
}
