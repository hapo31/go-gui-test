package calculator

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Calculator struct {
	ResultLabel     *widget.Label
	NumberButtons   []*widget.Button
	OperatorButtons []*widget.Button
	Buffer          string
	Result          int
	numberChannel   chan string
	operatorChannel chan string
}

type CalculatorWindow struct {
	fyne.Window
	*Calculator
}

func NewCalculatorWindow(app fyne.App) *CalculatorWindow {
	w := app.NewWindow("g-calc")
	calc := NewCalculator()
	w.Resize(fyne.NewSize(200, 30))
	var o []fyne.CanvasObject
	for _, v := range calc.NumberButtons {
		o = append(o, v)
	}
	for _, v := range calc.OperatorButtons {
		o = append(o, v)
	}

	c := container.NewWithoutLayout(o...)
	w.SetContent(c)

	return &CalculatorWindow{
		Window: w, Calculator: calc,
	}
}

func NewCalculator() *Calculator {

	calc := &Calculator{}

	calc.numberChannel = make(chan string)
	calc.operatorChannel = make(chan string)

	for i := 0; i < 10; i = i + 1 {
		calc.NumberButtons = append(calc.NumberButtons, widget.NewButton(strconv.Itoa(i), func() {
			calc.numberChannel <- strconv.Itoa(i)
		}))
	}

	for _, v := range []string{"C", "+", "-", "/", "*", "="} {
		calc.OperatorButtons = append(calc.OperatorButtons, widget.NewButton(v, func() {
			calc.operatorChannel <- v
		}))
	}

	return calc
}
