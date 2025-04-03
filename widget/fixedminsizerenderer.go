package widget

import (
	"fyne.io/fyne/v2"
)

type FixedMinSizeRenderer struct {
	Object       fyne.CanvasObject
	FixedMinSize fyne.Size
}

var _ fyne.WidgetRenderer = (*FixedMinSizeRenderer)(nil)

// NewFixedMinSizeRenderer creates a new renderer for the given object with fixed minimal size.
func NewFixedMinSizeRenderer(object fyne.CanvasObject, fixedSize fyne.Size) *FixedMinSizeRenderer {
	return &FixedMinSizeRenderer{
		Object:       object,
		FixedMinSize: fixedSize,
	}
}

func (f *FixedMinSizeRenderer) Layout(size fyne.Size) {
	f.Object.Resize(size)
}

func (f *FixedMinSizeRenderer) MinSize() fyne.Size {
	var minSize fyne.Size

	minSize = fyne.Size{
		Width:  0,
		Height: 0,
	}

	if f.FixedMinSize.Width > 0 {
		minSize.Width = f.FixedMinSize.Width
	}

	if f.FixedMinSize.Height > 0 {
		minSize.Height = f.FixedMinSize.Height
	}

	return minSize
}

func (f *FixedMinSizeRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{f.Object}
}

func (f *FixedMinSizeRenderer) Refresh() {
}

func (f *FixedMinSizeRenderer) Destroy() {
}
