package main

import (
	"fyne.io/fyne/v2/app"

	"github.com/Bjornsrud/GoGuiCalc/internal/gui"
)

func main() {
	a := app.New()
	w := gui.NewCalculatorWindow(a)
	w.ShowAndRun()
}
