package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewCalculatorWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Go GUI Calc")

	display := widget.NewLabel("0")

	w.SetContent(container.NewVBox(display))

	return w
}