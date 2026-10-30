package business

import "tic-tack-toe/internal/logger"

type Calc struct {
	log *logger.Logger
}

func NewCalc(l *logger.Logger) *Calc {
	return &Calc{log: l}
}

func (c *Calc) Add(a, b int) int {
	c.log.Log("Added")
	return a + b
}
