package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/halsten-dev/komok/manager"
	"math"
	"strconv"
)

type Entry struct {
	widget.Entry
	tabManagement    bool
	selectAllOnFocus bool
	shortcutsManager *manager.ShortcutsManager
	onlyNumerical    bool
	round            bool
}

func (e *Entry) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		e.Entry.TypedShortcut(s)
		return
	}

	e.shortcutsManager.TriggerShortcut(s)
}

func (e *Entry) AcceptsTab() bool {
	return e.MultiLine && e.tabManagement
}

func (e *Entry) FocusGained() {
	e.Entry.FocusGained()
	if e.selectAllOnFocus {
		e.SelectAll()
	}
}

// SelectAll selects all the text in the entry.
func (e *Entry) SelectAll() {
	e.TypedShortcut(&fyne.ShortcutSelectAll{})
}

func (e *Entry) SetSelectAllOnFocus(b bool) {
	e.selectAllOnFocus = b
}

func NewEntry(sm *manager.ShortcutsManager) *Entry {
	e := &Entry{}
	e.ExtendBaseWidget(e)

	e.shortcutsManager = sm
	e.Wrapping = fyne.TextWrap(fyne.TextTruncateClip)
	e.onlyNumerical = false
	e.round = false
	e.tabManagement = false

	return e
}

func NewMultilineEntry(sm *manager.ShortcutsManager, tm bool, wrap bool) *Entry {
	e := &Entry{}
	e.ExtendBaseWidget(e)

	e.shortcutsManager = sm
	e.MultiLine = true
	e.tabManagement = tm
	e.onlyNumerical = false
	e.round = false

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
	e.shortcutsManager = sm
	e.MultiLine = false
	e.tabManagement = false
	e.onlyNumerical = true
	e.round = round

	e.Validator = func(s string) error {
		if len(s) > 0 {
			if e.onlyNumerical {
				numericalValue, _ = strconv.ParseFloat(e.Text, 64)

				if e.round {
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
