package calculator

import "testing"

func TestPressDigitBuildsNumberInDisplay(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(1)
	calc.PressDigit(2)
	calc.PressDigit(3)

	got := calc.Display()
	want := "123"

	if got != want {
		t.Fatalf("Display() = %q, want %q", got, want)
	}
}

func TestInitialDisplayIsZero(t *testing.T) {
	calc := NewCalculator()

	got := calc.Display()
	want := "0"

	if got != want {
		t.Fatalf("initial Display() = %q, want %q", got, want)
	}
}

func TestFirstDigitReplacesZero(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(5)

	got := calc.Display()
	want := "5"

	if got != want {
		t.Fatalf("Display() after first digit = %q, want %q", got, want)
	}
}


func TestPressDotAddsDecimalPoint(t *testing.T) {
	calc := NewCalculator()

	calc.PressDot()

	got := calc.Display()
	want := "0."

	if got != want {
		t.Fatalf("Display() after PressDot = %q, want %q", got, want)
	}
}

func TestPressDotOnlyAllowedOnce(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(1)
	calc.PressDot()
	calc.PressDigit(5)
	calc.PressDot() // second dot should do nothing

	got := calc.Display()
	want := "1.5"

	if got != want {
		t.Fatalf("Display() with second dot = %q, want %q", got, want)
	}
}
