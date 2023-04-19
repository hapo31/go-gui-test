package calculator

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CalculatorElement struct {
	Calclator
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
	*CalculatorElement
}

func NewCalculatorWindow(app fyne.App) *CalculatorWindow {
	w := app.NewWindow("g-calc")
	calc := NewCalculatorElement()
	w.Resize(fyne.NewSize(200, 30))
	var n []fyne.CanvasObject
	var o []fyne.CanvasObject
	for _, v := range calc.NumberButtons {
		n = append(n, v)
	}
	for _, v := range calc.OperatorButtons {
		o = append(o, v)
	}

	c := container.NewVBox(
		calc.ResultLabel,
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(3, n...),
			container.NewGridWithColumns(3, o...),
		))
	w.SetContent(c)

	return &CalculatorWindow{
		Window: w, CalculatorElement: calc,
	}
}

func NewCalculatorElement() *CalculatorElement {

	calc := &CalculatorElement{}

	calc.numberChannel = make(chan string)
	calc.operatorChannel = make(chan string)

	calc.Calclator = NewCalculator()

	calc.ResultLabel = widget.NewLabel("test")

	for i := 0; i < 10; i += 1 {
		number := strconv.Itoa(i)
		calc.NumberButtons = append(calc.NumberButtons, widget.NewButton(number, func() {
			calc.numberChannel <- number
		}))
	}

	for _, v := range []string{"C", "+", "-", "/", "*", "="} {
		operator := v
		calc.OperatorButtons = append(calc.OperatorButtons, widget.NewButton(operator, func() {
			calc.operatorChannel <- operator
		}))
	}

	go func() {
		for {
			select {
			case n := <-calc.numberChannel:
				calc.Calclator.AddStr(n)
			case o := <-calc.operatorChannel:
				calc.Calclator.AddStr(o)
			default:
				continue
			}
			calc.ResultLabel.SetText(calc.Calclator.Buffer)
		}
	}()

	return calc
}
