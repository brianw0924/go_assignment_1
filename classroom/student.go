package classroom

import (
	"fmt"
	"math/rand"
	"time"
)

type Student struct {
	Name     string
	Ready    chan struct{}
	Done     chan struct{}
	Answered bool
}

func NewStudent(name string) *Student {
	return &Student{
		Name:     name,
		Ready:    make(chan struct{}, 1),
		Done:     make(chan struct{}, 1),
		Answered: false,
	}
}

func (s *Student) Start(question chan *Question, answer chan *Answer) {
	for {
		s.WaitQuestion()
		s.SeeAndThink()
		s.AnswerQuestion(question, answer)
	}
}

func (s *Student) WaitQuestion() {
	<-s.Ready
}

func (s *Student) QuestionReady() {
	s.Ready <- struct{}{}
}

func (s *Student) SeeAndThink() {
	time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2001)))
}

func (s *Student) AnswerQuestion(question chan *Question, answer chan *Answer) {
	for {
		select {
		case q := <-question:
			s.Answered = true
			ans := s.CalcAnswer(q)
			ans.val += rand.Intn(5) - 2
			fmt.Printf("Student %s: %d %c %d = %d\n", s.Name, q.LeftOperand, q.Operator, q.RightOperand, ans.val)
			answer <- ans
			return
		case <-s.Done:
			return
		}
	}
}

func (s *Student) CalcAnswer(q *Question) *Answer {
	switch q.Operator {
	case '+':
		return NewAnswer(s.Name, q.LeftOperand+q.RightOperand)
	case '-':
		return NewAnswer(s.Name, q.LeftOperand-q.RightOperand)
	case '*':
		return NewAnswer(s.Name, q.LeftOperand*q.RightOperand)
	case '/':
		return NewAnswer(s.Name, q.LeftOperand/q.RightOperand)
	default:
		panic(fmt.Sprintf("Invalid operator: %c", q.Operator))
	}
}

func (s *Student) QuestionDone(winner string) {
	if winner != s.Name && winner != "" {
		fmt.Printf("Student %s: %s, you win.\n", s.Name, winner)
	}

	if s.Answered {
		s.Answered = false
	} else {
		s.Done <- struct{}{}
	}
}
