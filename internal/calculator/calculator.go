package calculator

type Calculator struct {
	display string
}

func NewCalculator() *Calculator {
	return &Calculator{
		display: "",
	}
}

func (c *Calculator) PressDigit(d int) {
	if d < 0 || d > 9 {
		return
	}

	c.display += string('0' + rune(d))
}

func (c *Calculator) Display() string {
	return c.display
}
