package widget

import (
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/halsten-dev/komok/manager"
)

type Entry struct {
	widget.Entry
	TabManagement    bool
	SelectAllOnFocus bool
	ShortcutsManager *manager.ShortcutsManager
	OnlyNumerical    bool
	Round            bool

	// OnTypedKey is called before the TypedKey event. Returns true if it was processed.
	OnTypedKey func(e *fyne.KeyEvent) bool
}

func (e *Entry) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		e.Entry.TypedShortcut(s)
		return
	}

	e.ShortcutsManager.TriggerShortcut(s)
}

func (e *Entry) AcceptsTab() bool {
	return e.MultiLine && e.TabManagement
}

func (e *Entry) FocusGained() {
	e.Entry.FocusGained()
	if e.SelectAllOnFocus {
		e.SelectAll()
	}
}

func (e *Entry) TypedKey(key *fyne.KeyEvent) {
	processed := false

	if e.OnTypedKey != nil {
		processed = e.OnTypedKey(key)
	}

	if !processed {
		e.Entry.TypedKey(key)
	}
}

// SelectAll selects all the text in the entry.
func (e *Entry) SelectAll() {
	if len(e.Text) > 0 {
		e.TypedShortcut(&fyne.ShortcutSelectAll{})
	}
}

func (e *Entry) SetSelectAllOnFocus(b bool) {
	e.SelectAllOnFocus = b
}

func NewEntry(sm *manager.ShortcutsManager) *Entry {
	e := &Entry{}
	e.ExtendBaseWidget(e)

	e.ShortcutsManager = sm
	e.Wrapping = fyne.TextWrap(fyne.TextTruncateClip)
	e.OnlyNumerical = false
	e.Round = false
	e.TabManagement = false

	return e
}

func NewMultilineEntry(sm *manager.ShortcutsManager, tm bool, wrap bool) *Entry {
	e := &Entry{}
	e.ExtendBaseWidget(e)

	e.ShortcutsManager = sm
	e.MultiLine = true
	e.TabManagement = tm
	e.OnlyNumerical = false
	e.Round = false

	if wrap {
		e.Wrapping = fyne.TextWrapWord
	} else {
		e.Wrapping = fyne.TextWrap(fyne.TextTruncateClip)
	}

	return e
}

func NewNumericalEntry(sm *manager.ShortcutsManager, round bool) *Entry {
	var numericalValue float64

	e := &Entry{}
	e.ExtendBaseWidget(e)
	e.ShortcutsManager = sm
	e.MultiLine = false
	e.TabManagement = false
	e.OnlyNumerical = true
	e.Round = round

	e.Validator = func(s string) error {
		if len(s) > 0 {
			if e.OnlyNumerical {
				numericalValue, _ = strconv.ParseFloat(e.Text, 64)

				if e.Round {
					e.Text = strconv.Itoa(int(math.Round(numericalValue)))
				} else {
					e.Text = strconv.FormatFloat(numericalValue, 'f', -1, 64)
				}
			}
		}

		return nil
	}

	return e
}
