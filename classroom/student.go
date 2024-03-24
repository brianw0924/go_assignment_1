package classroom

import (
	"fmt"
	"math/rand"
	"time"
)

type Student struct {
	Name string
}

func NewStudent(name string) *Student {
	return &Student{
		Name: name,
	}
}

func (s *Student) QuestionReady(question <-chan *Question, answer chan<- *Answer, winner <-chan *Answer) {
	s.SeeAndThink()
	go s.AnswerQuestion(question, answer)
	go s.WaitWinner(winner)
}

func (s *Student) WaitWinner(winner <-chan *Answer) {
	ans := <-winner
	fmt.Printf("Student %s: %s, Q%d you win.\n", s.Name, ans.fromStudent, ans.QuestionId)
}

func (s *Student) SeeAndThink() {
	time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(5001)))
}

func (s *Student) AnswerQuestion(question <-chan *Question, answer chan<- *Answer) {
	if q, ok := <-question; ok {
		ans := s.calcAnswer(q)
		fmt.Printf("Student %s: Q%d: %d %c %d = %d!\n", s.Name, q.Id, q.LeftOperand, q.Operator, q.RightOperand, ans.val)
		answer <- ans
		close(answer)
	}
}

func (s *Student) calcAnswer(q *Question) *Answer {
	switch q.Operator {
	case '+':
		return NewAnswer(s.Name, q.LeftOperand+q.RightOperand, q.Id)
	case '-':
		return NewAnswer(s.Name, q.LeftOperand-q.RightOperand, q.Id)
	case '*':
		return NewAnswer(s.Name, q.LeftOperand*q.RightOperand, q.Id)
	case '/':
		return NewAnswer(s.Name, q.LeftOperand/q.RightOperand, q.Id)
	default:
		panic(fmt.Sprintf("Invalid operator: %c", q.Operator))
	}
}
