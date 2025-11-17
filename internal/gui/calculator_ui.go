package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Bjornsrud/GoGuiCalc/internal/calculator"
)

func NewCalculatorWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("GoCalculate")
	w.Resize(fyne.NewSize(240, 320))


	display := widget.NewLabel("0")
	calc := calculator.NewCalculator()

	makeDigitButton := func(n int) *widget.Button {
    return widget.NewButton(
        fmt.Sprintf("%d", n),
        func() {
            calc.PressDigit(n)
            display.SetText(calc.Display())
        },
    )
}

	buttons := container.NewGridWithRows(4,
		container.NewGridWithColumns(3,
			makeDigitButton(1),
			makeDigitButton(2),
			makeDigitButton(3),
		),
		container.NewGridWithColumns(3,
			makeDigitButton(4),
			makeDigitButton(5),
			makeDigitButton(6),
		),
		container.NewGridWithColumns(3,
			makeDigitButton(7),
			makeDigitButton(8),
			makeDigitButton(9),
		),
		container.NewGridWithColumns(3,
			makeDigitButton(0),
			widget.NewLabel(""), // tom plass
			widget.NewLabel(""), // tom plass
		),
	)

	w.SetContent(container.NewVBox(
		display,
		buttons,
	))

	return w
}

