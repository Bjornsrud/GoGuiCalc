package calculator

import (
	"strconv"
	"strings"
)

type Calculator struct {
	display string
	accumulator float64
	operator string
	overwrite bool
}

func NewCalculator() *Calculator {
	return &Calculator{
		display: "0",
	}
}

func (c *Calculator) PressDigit(d int) {
	if d < 0 || d > 9 {
		return
	}

	digit := string('0' + rune(d))

	if c.display == "0" {
		c.display = digit
		return
	}

	c.display += digit
}

func (c *Calculator) PressDot() {
	if strings.Contains(c.display, ".") {
		return
	}

	c.display += "."
}

func (c *Calculator) PressClear() { 
	c.display = "0"
}


func (c *Calculator) Display() string {
	return c.display
}

func (c *Calculator) Value() float64 {
    v, err := strconv.ParseFloat(c.display, 64)
	if err != nil {
		return 0
	}
	return v
}

func (c *Calculator) PressOperator(o string) {
	return
}
