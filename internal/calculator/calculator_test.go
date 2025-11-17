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

func TestPressDigitAfterOperatorOverwritesDisplay(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(4)
	calc.PressDigit(2)
	calc.PressOperator("+")

	calc.PressDigit(3)

	got := calc.Display()
	want := "3"

	if got != want {
		t.Fatalf("Display() after operator and digit = %q, want %q", got, want)
	}
}

func TestPressOperatorEvaluatesPreviousOperator(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(2)
	calc.PressOperator("+")
	calc.PressDigit(3)
	calc.PressOperator("+")

	gotAcc := calc.accumulator
	wantAcc := 5.0

	if gotAcc != wantAcc {
		t.Fatalf("accumulator after evaluating 2 + 3 = %v, want %v", gotAcc, wantAcc)
	}

	if calc.operator != "+" {
		t.Fatalf("operator = %q, want %q", calc.operator, "+")
	}

	if calc.Display() != "5" {
		t.Fatalf("display = %q, want %q", calc.Display(), "5")
	}
}

func TestPressEqualsEvaluatesPendingOperation(t *testing.T) {
	calc := NewCalculator()

	// 2 + 3 =
	calc.PressDigit(2)
	calc.PressOperator("+")
	calc.PressDigit(3)
	calc.PressEquals()

	got := calc.Display()
	want := "5"

	if got != want {
		t.Fatalf("Display() after 2 + 3 = should be %q, got %q", want, got)
	}
}

func TestPressEqualsClearsOperatorAndSetsOverwrite(t *testing.T) {
	calc := NewCalculator()

	// 2 + 3 =
	calc.PressDigit(2)
	calc.PressOperator("+")
	calc.PressDigit(3)
	calc.PressEquals()

	if calc.operator != "" {
		t.Fatalf("operator after Equals = %q, want empty", calc.operator)
	}

	calc.PressDigit(4)

	got := calc.Display()
	want := "4"

	if got != want {
		t.Fatalf("Display() after Equals and then digit = %q, want %q", got, want)
	}
}

func TestPressEqualsCanBeRepeatedForAddition(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(2)
	calc.PressOperator("+")
	calc.PressDigit(3)
	calc.PressEquals()

	if calc.Display() != "5" {
		t.Fatalf("after first Equals, display = %q, want %q", calc.Display(), "5")
	}


	calc.PressEquals()
	if calc.Display() != "8" {
		t.Fatalf("after second Equals, display = %q, want %q", calc.Display(), "8")
	}

	// and another =
	calc.PressEquals()
	if calc.Display() != "11" {
		t.Fatalf("after third Equals, display = %q, want %q", calc.Display(), "11")
	}
}

func TestPressEqualsCanBeRepeatedForSubtraction(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(2)
	calc.PressOperator("-")
	calc.PressDigit(3)
	calc.PressEquals()

	if calc.Display() != "-1" {
		t.Fatalf("after first Equals, display = %q, want %q", calc.Display(), "-1")
	}

	calc.PressEquals()
	if calc.Display() != "-4" {
		t.Fatalf("after second Equals, display = %q, want %q", calc.Display(), "-4")
	}

	calc.PressEquals()
	if calc.Display() != "-7" {
		t.Fatalf("after third Equals, display = %q, want %q", calc.Display(), "-7")
	}
}

func TestPressEqualsCanBeRepeatedForMultiplication(t *testing.T) {
	calc := NewCalculator()

	// 2 * 3 =
	calc.PressDigit(2)
	calc.PressOperator("*")
	calc.PressDigit(3)
	calc.PressEquals()

	if calc.Display() != "6" {
		t.Fatalf("after first Equals, display = %q, want %q", calc.Display(), "6")
	}

	// another =
	calc.PressEquals()
	if calc.Display() != "18" {
		t.Fatalf("after second Equals, display = %q, want %q", calc.Display(), "18")
	}

	// and another =
	calc.PressEquals()
	if calc.Display() != "54" {
		t.Fatalf("after third Equals, display = %q, want %q", calc.Display(), "54")
	}
}

func TestPressEqualsCanBeRepeatedForDivision(t *testing.T) {
	calc := NewCalculator()

	// 100 / 2 =
	calc.PressDigit(1)
	calc.PressDigit(0)
	calc.PressDigit(0)
	calc.PressOperator("/")
	calc.PressDigit(2)
	calc.PressEquals()

	if calc.Display() != "50" {
		t.Fatalf("after first Equals, display = %q, want %q", calc.Display(), "50")
	}

	calc.PressEquals()
	if calc.Display() != "25" {
		t.Fatalf("after second Equals, display = %q, want %q", calc.Display(), "25")
	}

	calc.PressEquals()
	if calc.Display() != "12.5" {
		t.Fatalf("after third Equals, display = %q, want %q", calc.Display(), "12.5")
	}
}

func TestDivisionByZeroShowsError(t *testing.T) {
	calc := NewCalculator()

	calc.PressDigit(5)
	calc.PressOperator("/")
	calc.PressDigit(0)
	calc.PressEquals()

	if calc.Display() != "Error" {
		t.Fatalf("division by zero: display = %q, want %q", calc.Display(), "Error")
	}
}




