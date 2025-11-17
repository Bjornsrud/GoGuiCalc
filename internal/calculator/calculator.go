package calculator

import (
	"strconv"
	"strings"
)

type Calculator struct {
	display      string
	accumulator  float64
	operator     string
	overwrite    bool
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
	c.accumulator = 0
	c.operator = ""
	c.lastOperator = ""
	c.lastOperand = 0
	c.overwrite = false
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

	if !c.overwrite {
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
				if current == 0 {
					c.display = "Error"
					c.operator = ""
					c.lastOperator = ""
					c.overwrite = true
					return
				}
				c.accumulator /= current
			}
		}

		c.lastOperand = current
		if c.operator != "" {
			c.lastOperator = c.operator
		} else {
			c.lastOperator = o
		}
	} else {
		if c.operator != "" && c.lastOperator != "" {
			switch c.lastOperator {
			case "+":
				c.accumulator += c.lastOperand
			case "-":
				c.accumulator -= c.lastOperand
			case "*":
				c.accumulator *= c.lastOperand
			case "/":
				if c.lastOperand == 0 {
					c.display = "Error"
					c.operator = ""
					c.lastOperator = ""
					c.overwrite = true
					return
				}
				c.accumulator /= c.lastOperand
			}
		}
	}

	c.display = strconv.FormatFloat(c.accumulator, 'f', -1, 64)
	c.operator = o
	c.overwrite = true
}

func (c *Calculator) PressEquals() {
	current := c.Value()

	if c.operator != "" {
		switch c.operator {
		case "+":
			c.accumulator += current
		case "-":
			c.accumulator -= current
		case "*":
			c.accumulator *= current
		case "/":
			if current == 0 {
				c.display = "Error"
				c.operator = ""
				c.lastOperator = ""
				c.overwrite = true
				return
			}
			c.accumulator /= current
		}
		c.lastOperator = c.operator
		c.lastOperand = current
		c.operator = ""
	} else if c.lastOperator != "" {
		c.accumulator = current

		switch c.lastOperator {
		case "+":
			c.accumulator += c.lastOperand
		case "-":
			c.accumulator -= c.lastOperand
		case "*":
			c.accumulator *= c.lastOperand
		case "/":
			if c.lastOperand == 0 {
				c.display = "Error"
				c.operator = ""
				c.lastOperator = ""
				c.overwrite = true
				return
			}
			c.accumulator /= c.lastOperand
		}
	} else {

		return
	}

	c.display = strconv.FormatFloat(c.accumulator, 'f', -1, 64)
	c.overwrite = true
}
