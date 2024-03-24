package classroom

import (
	"math/rand"
)

var operator = []rune{'+', '-', '*', '/'}

type Question struct {
	Id           int
	LeftOperand  int
	RightOperand int
	Operator     rune
}

func NewQuestion(id int) *Question {
	return &Question{
		Id:           id,
		LeftOperand:  rand.Intn(101),
		RightOperand: rand.Intn(101),
		Operator:     operator[rand.Intn(len(operator))],
	}
}
