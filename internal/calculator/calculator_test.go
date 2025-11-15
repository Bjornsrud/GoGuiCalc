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
