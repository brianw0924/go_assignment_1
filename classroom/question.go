package classroom

import (
	"math/rand"
)

var operator = []rune{'+', '-', '*', '/'}

type Question struct {
	LeftOperand  int
	RightOperand int
	Operator     rune
}

func NewQuestion() *Question {
	return &Question{
		LeftOperand:  rand.Intn(101),
		RightOperand: rand.Intn(101),
		Operator:     operator[rand.Intn(len(operator))],
	}
}
