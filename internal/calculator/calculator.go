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
	lastOperator string
	lastOperand  float64
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


	if c.overwrite {
		c.display = digit
		c.overwrite = false
		return
	}

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
	current := c.Value()

	if c.operator == "" {
		c.accumulator = current
	} else {
		switch c.operator {
		case "+":
			c.accumulator += current
		case "-":
			c.accumulator -= current
		case "*":
			c.accumulator *= current
		case "/":
			c.accumulator /= current
		}
	}

	

	c.display = strconv.FormatFloat(c.accumulator, 'f', -1, 64)
	c.operator = o
	c.overwrite = true
}

func (c *Calculator) PressEquals() {
	current := c.Value()

	// Første gang vi trykker "=", bruker vi pending operator
	if c.operator != "" {
		switch c.operator {
		case "+":
			c.accumulator += current
		case "-":
			c.accumulator -= current
		case "*":
			c.accumulator *= current
		case "/":
			c.accumulator /= current
		}

		// Husk hva vi gjorde, slik at vi kan gjenta det
		c.lastOperator = c.operator
		c.lastOperand = current

		c.operator = ""
	} else if c.lastOperator != "" {
		// Ingen pending operator, men vi har en tidligere "=" operasjon
		c.accumulator = current

		switch c.lastOperator {
		case "+":
			c.accumulator += c.lastOperand
		case "-":
			c.accumulator -= c.lastOperand
		case "*":
			c.accumulator *= c.lastOperand
		case "/":
			c.accumulator /= c.lastOperand
		}
	} else {
		// Ingen operatør og ingen historikk: gjør ingenting
		return
	}

	c.display = strconv.FormatFloat(c.accumulator, 'f', -1, 64)
	c.overwrite = true
}

