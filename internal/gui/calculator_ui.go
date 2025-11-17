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
	w.Resize(fyne.NewSize(285, 240))

	display := canvas.NewText("0", color.White)
	display.TextSize = 32
	display.Alignment = fyne.TextAlignTrailing

	calc := calculator.NewCalculator()

	makeColorButton := func(label string, bg color.Color, onTap func()) *fyne.Container {
		txt := canvas.NewText(label, color.White)
		txt.Alignment = fyne.TextAlignCenter
		txt.TextSize = 18

		rect := canvas.NewRectangle(bg)

		btn := widget.NewButton("", onTap)
		btn.Importance = widget.LowImportance

		return container.NewStack(rect, btn, container.NewCenter(txt))
	}

	digit := func(n int) *fyne.Container {
		return makeColorButton(
			fmt.Sprintf("%d", n),
			color.RGBA{50, 50, 50, 255},
			func() {
				calc.PressDigit(n)
				display.Text = calc.Display()
				display.Refresh()
			},
		)
	}

	op := func(label string, opSymbol string) *fyne.Container {
		return makeColorButton(
			label,
			color.RGBA{70, 70, 70, 255},
			func() {
				calc.PressOperator(opSymbol)
				display.Text = calc.Display()
				display.Refresh()
			},
		)
	}

	dotButton := makeColorButton(
		".",
		color.RGBA{50, 50, 50, 255},
		func() {
			calc.PressDot()
			display.Text = calc.Display()
			display.Refresh()
		},
	)

	equalsButton := makeColorButton(
		"=",
		color.RGBA{60, 100, 60, 255},
		func() {
			calc.PressEquals()
			display.Text = calc.Display()
			display.Refresh()
		},
	)

	clearButton := makeColorButton(
		"Clear",
		color.RGBA{140, 80, 10, 255},
		func() {
			calc.PressClear()
			display.Text = calc.Display()
			display.Refresh()
		},
	)

	aboutButton := makeColorButton(
		"About",
		color.RGBA{90, 90, 90, 255},
		func() {
			display.Text = "Bjornsrud@github"
			display.Refresh()
		},
	)

	buttons := container.NewGridWithColumns(4,
		digit(7),
		digit(8),
		digit(9),
		op("/", "/"),

		digit(4),
		digit(5),
		digit(6),
		op("x", "*"),

		digit(1),
		digit(2),
		digit(3),
		op("-", "-"),

		digit(0),
		dotButton,
		equalsButton,
		op("+", "+"),
	)

	headerRow := container.NewGridWithColumns(2,
		clearButton,
		aboutButton,
	)

	w.SetContent(container.NewVBox(
		display,
		headerRow,
		buttons,
	))
	return w
}
