package calculator

import (
	"fyne.io/fyne/v2/widget"
)

type Calculator struct {
	ResultLabel     *widget.Label
	NumberButtons   []*widget.Button
	OperatorButtons []*widget.Button
	Buffer          string
	Result          int
}
