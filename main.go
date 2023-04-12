package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/hapo31/calculator"
)

func main() {
	a := app.New()
	w := calculator.NewCalculatorWindow(a)

	w.Resize(fyne.NewSize(480, 360))
	w.ShowAndRun()

}
