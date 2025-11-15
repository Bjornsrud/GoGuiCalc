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

func TestPressClearResetsDisplayToZero(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(4)
	calc.PressDigit(2)
	calc.PressClear()

	got := calc.Display()
	want := "0"

	if got != want {
		t.Fatalf("PressClear should reset display to %q, got %q", want, got)
	}
}

func TestPressClearOnZeroLeavesZero(t *testing.T) {
	calc := NewCalculator()

	calc.PressClear()

	got := calc.Display()
	want := "0"

	if got != want {
		t.Fatalf("PressClear on initial zero should keep display as %q, got %q", want, got)
	}
}

func TestValueParsesIntegerDisplay(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(1)
	calc.PressDigit(2)

	got := calc.Value()
	want := 12.0

	if got != want {
		t.Fatalf("Value() = %v, want %v", got, want)
	}
}

func TestValueParsesDecimalDisplay(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(3)
	calc.PressDot()
	calc.PressDigit(5)

	got := calc.Value()
	want := 3.5

	if got != want {
		t.Fatalf("Value() = %v, want %v", got, want)
	}
}

func TestValueParsesZero(t *testing.T) {
	calc := NewCalculator()

	got := calc.Value()
	want := 0.0

	if got != want {
		t.Fatalf("Value() for zero = %v, want %v", got, want)
	}
}

func TestPressOperatorStoresAccumulatorAndOperator(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(4)
	calc.PressDigit(2)

	calc.PressOperator("+")

	if calc.accumulator != 42 {
		t.Fatalf("accumulator = %v, want 42", calc.accumulator)
	}

	if calc.operator != "+" {
		t.Fatalf("operator = %q, want \"+\"", calc.operator)
	}

	if !calc.overwrite {
		t.Fatalf("overwrite should be true after pressing operator")
	}
}
