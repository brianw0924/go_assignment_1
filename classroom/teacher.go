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
		s.QuestionDone(winner)
	}
}

func (t *Teacher) Start() {
	for {
		t.WarmUp()
		q, ans := t.AskQuestion()
		t.QuestionReady()
		winner := t.WaitAnswer(q, ans)
		t.QuestionDone(winner)
	}
}

func (t *Teacher) WarmUp() {
	fmt.Println("Teacher: Guys, are you ready?")
	time.Sleep(time.Second * 3)
}

func (t *Teacher) AskQuestion() (*Question, *Answer) {
	q := NewQuestion()
	fmt.Printf("Teacher: %d %c %d = ?\n", q.LeftOperand, q.Operator, q.RightOperand)
	t.QuestionChan <- q
	return q, t.CalcAnswer("teacher", q)
}

func (t *Teacher) WaitAnswer(q *Question, ans *Answer) string {

	for i := 0; i < len(t.StudentList); i += 1 {
		ansFromStudent := <-t.AnswerChan
		if res := t.VerifyAnswer(ans, ansFromStudent); res {
			winner := ansFromStudent.fromStudent
			fmt.Printf("Teacher: %s, you are right!\n", winner)
			return winner
		} else {
			fmt.Printf("Teacher: %s, you are wrong!\n", ansFromStudent.fromStudent)
			t.QuestionChan <- q
		}
	}
	// All students have wrong answer. Remove the last question
	<-t.QuestionChan

	fmt.Printf("Teacher: Boooo~ Answer is %d\n", ans.val)

	return ""
}

func (t *Teacher) VerifyAnswer(a1 *Answer, a2 *Answer) bool {
	return a1.val == a2.val
}

func (t *Teacher) CalcAnswer(name string, q *Question) *Answer {
	switch q.Operator {
	case '+':
		return NewAnswer(name, q.LeftOperand+q.RightOperand)
	case '-':
		return NewAnswer(name, q.LeftOperand-q.RightOperand)
	case '*':
		return NewAnswer(name, q.LeftOperand*q.RightOperand)
	case '/':
		return NewAnswer(name, q.LeftOperand/q.RightOperand)
	default:
		panic(fmt.Sprintf("Invalid operator: %c", q.Operator))
	}
}
