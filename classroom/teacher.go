package classroom

import (
	"fmt"
	"time"
)

type Teacher struct {
	StudentList  []*Student
	QuestionChan chan *Question
	AnswerChan   chan *Answer
}

func NewTeacher() *Teacher {
	return &Teacher{
		StudentList:  make([]*Student, 0),
		QuestionChan: make(chan *Question, 1),
		AnswerChan:   make(chan *Answer, 1),
	}
}

func (t *Teacher) Register(s *Student) {
	t.StudentList = append(t.StudentList, s)
}

func (t *Teacher) QuestionReady() {
	for _, s := range t.StudentList {
		s.QuestionReady()
	}
}

func (t *Teacher) QuestionDone(winner string) {
	for _, s := range t.StudentList {
		if s.Name != winner {
			s.QuestionDone(winner)
		}
	}
}

func (t *Teacher) Start() {
	for {
		t.WarmUp()
		t.AskQuestion()
		t.QuestionReady()
		winner := t.WaitAnswer()
		t.QuestionDone(winner)

	}
}

func (t *Teacher) WarmUp() {
	fmt.Println("Teacher: Guys, are you ready?")
	time.Sleep(time.Second * 3)
}

func (t *Teacher) AskQuestion() {
	q := NewQuestion()
	fmt.Printf("Teacher: %d %c %d = ?\n", q.LeftOperand, q.Operator, q.RightOperand)
	t.QuestionChan <- q
}
func (t *Teacher) WaitAnswer() (winner string) {
	ans := <-t.AnswerChan
	winner = ans.fromStudent
	fmt.Printf("Teacher: %s, you are right!\n", winner)
	return winner
}
