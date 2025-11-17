package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Bjornsrud/GoGuiCalc/internal/calculator"
)

func NewCalculatorWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("GoCalculate")
	w.Resize(fyne.NewSize(250, 240))

	// Display som canvas.Text, så vi kan styre størrelse og farge
	display := canvas.NewText("0", color.White)
	display.TextSize = 32
	display.Alignment = fyne.TextAlignTrailing

	calc := calculator.NewCalculator()

	// Generell factory for fargede knapper
	makeColorButton := func(label string, bg color.Color, onTap func()) *fyne.Container {
		txt := canvas.NewText(label, color.White)
		txt.Alignment = fyne.TextAlignCenter
		txt.TextSize = 18

		rect := canvas.NewRectangle(bg)

		btn := widget.NewButton("", onTap)
		btn.Importance = widget.LowImportance

		return container.NewMax(rect, btn, container.NewCenter(txt))
	}

	// Tallknapper
	digit := func(n int) *fyne.Container {
		return makeColorButton(
			fmt.Sprintf("%d", n),
			color.RGBA{50, 50, 50, 255}, // mørk grå
			func() {
				calc.PressDigit(n)
				display.Text = calc.Display()
				display.Refresh()
			},
		)
	}

	// Operator-knapper
	op := func(symbol string) *fyne.Container {
		return makeColorButton(
			symbol,
			color.RGBA{70, 70, 70, 255}, // Lysere grå
			func() {
				calc.PressOperator(symbol)
				display.Text = calc.Display()
				display.Refresh()
			},
		)
	}

	// Spesialknapper
	dotButton := makeColorButton(
		".",
		color.RGBA{50, 50, 50, 255}, // samme som tall
		func() {
			calc.PressDot()
			display.Text = calc.Display()
			display.Refresh()
		},
	)

	equalsButton := makeColorButton(
		"=",
		color.RGBA{60, 100, 60, 255}, // grønn
		func() {
			calc.PressEquals()
			display.Text = calc.Display()
			display.Refresh()
		},
	)

	clearButton := makeColorButton(
		"Clear",
		color.RGBA{140, 80, 10, 255}, // orange
		func() {
			calc.PressClear()
			display.Text = calc.Display()
			display.Refresh()
		},
	)

	// 4x4 grid med alle knapper
	buttons := container.NewGridWithColumns(4,
		// Rad 1
		digit(7),
		digit(8),
		digit(9),
		op("/"),

		// Rad 2
		digit(4),
		digit(5),
		digit(6),
		op("x"),

		// Rad 3
		digit(1),
		digit(2),
		digit(3),
		op("-"),

		// Rad 4
		digit(0),
		dotButton,
		equalsButton,
		op("+"),
	)

	w.SetContent(container.NewVBox(
		display,
		clearButton,
		buttons,
	))

	return w
}
