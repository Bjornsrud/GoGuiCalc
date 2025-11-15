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
