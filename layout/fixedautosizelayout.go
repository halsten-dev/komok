package layout

import (
	"log"

	"fyne.io/fyne/v2"
)

type FixedAutoSize struct {
	fixedSize  float32
	padding    float32
	horizontal bool
	reversed   bool
}

func NewFixedAutoSizeLayout(fixedSize float32, horizontal bool) fyne.Layout {
	return &FixedAutoSize{fixedSize, 0, horizontal, false}
}

func NewFixedAutoSizeLayoutPadded(fixedSize float32, padding float32, horizontal bool) fyne.Layout {
	return &FixedAutoSize{fixedSize, padding, horizontal, false}
}

func NewFixedAutoSizeLayoutReversed(fixedSize float32, horizontal bool) fyne.Layout {
	return &FixedAutoSize{fixedSize, 0, horizontal, true}
}

func NewFixedAutoSizeLayoutPaddedReversed(fixedSize float32, padding float32, horizontal bool) fyne.Layout {
	return &FixedAutoSize{fixedSize, padding, horizontal, true}
}

func (fas *FixedAutoSize) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		// minSize = minSize.Max(child.MinSize())
		
		if minSize.Width < child.MinSize().Width {
			minSize.Width = child.MinSize().Width
		}

		if minSize.Height < child.MinSize().Height {
			minSize.Height = child.MinSize().Height
		}
	}

	if fas.horizontal {
		minSize.Width += fas.fixedSize + fas.padding
	} else {
		minSize.Height += fas.fixedSize + fas.padding
	}

	return minSize
}

func (fas *FixedAutoSize) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	topLeft := fyne.NewPos(0, 0)

	if len(objects) > 2 || len(objects) < 2 {
		log.Println("Misused FixedAutoSizeLayout !")
		return
	}

	object1 := objects[0]
	object2 := objects[1]

	object1.Move(topLeft)

	if !fas.reversed {
		if fas.horizontal {
			object1.Resize(fyne.NewSize(fas.fixedSize, containerSize.Height))

			object2.Move(topLeft.AddXY(object1.Size().Width+fas.padding, 0))
			object2.Resize(fyne.NewSize(containerSize.Width-fas.fixedSize-fas.padding, containerSize.Height))
		} else {
			object1.Resize(fyne.NewSize(containerSize.Width, fas.fixedSize))

			object2.Move(topLeft.AddXY(0, object1.Size().Height+fas.padding))
			object2.Resize(fyne.NewSize(containerSize.Width, containerSize.Height-fas.fixedSize-fas.padding))
		}
	} else {
		if fas.horizontal {
			object1.Resize(fyne.NewSize(containerSize.Width-fas.fixedSize-fas.padding, containerSize.Height))

			object2.Move(topLeft.AddXY(object1.Size().Width+fas.padding, 0))
			object2.Resize(fyne.NewSize(fas.fixedSize, containerSize.Height))
		} else {
			object1.Resize(fyne.NewSize(containerSize.Width, containerSize.Height-fas.fixedSize-fas.padding))

			object2.Move(topLeft.AddXY(0, object1.Size().Height+fas.padding))
			object2.Resize(fyne.NewSize(containerSize.Width, fas.fixedSize))
		}
	}

}
