package classroom

import (
	"fmt"
	"time"
)

type Teacher struct {
	StudentList []*Student
}

func NewTeacher() *Teacher {
	return &Teacher{
		StudentList: make([]*Student, 0),
	}
}

func (t *Teacher) Register(s *Student) {
	t.StudentList = append(t.StudentList, s)
}

func (t *Teacher) QuestionReady(question <-chan *Question, answer chan<- *Answer, winner <-chan *Answer) {
	for _, s := range t.StudentList {
		go s.QuestionReady(question, answer, winner)
	}
}

func (t *Teacher) QuestionDone(winner chan<- *Answer, ans *Answer) {
	for i := 0; i < len(t.StudentList); i += 1 {
		winner <- ans
	}
}

func (t *Teacher) Start() {
	t.WarmUp()
	qid := 0
	for {
		time.Sleep(time.Second * 1)

		go func() {
			question := make(chan *Question, 1)
			answer := make(chan *Answer, 1)
			winner := make(chan *Answer, len(t.StudentList))

			t.AskQuestion(qid, question)
			t.QuestionReady(question, answer, winner)
			ans := t.WaitAnswer(answer)
			t.QuestionDone(winner, ans)
		}()

		qid += 1
	}
}

func (t *Teacher) WarmUp() {
	fmt.Println("Teacher: Guys, are you ready?")
	time.Sleep(time.Second * 3)
}

func (t *Teacher) AskQuestion(qid int, question chan<- *Question) {
	q := NewQuestion(qid)
	fmt.Printf("Teacher: Q%d: %d %c %d = ?\n", q.Id, q.LeftOperand, q.Operator, q.RightOperand)
	question <- q
	close(question)
}
func (t *Teacher) WaitAnswer(answer <-chan *Answer) *Answer {
	ans := <-answer
	fmt.Printf("Teacher: %s, Q%d you are right!\n", ans.fromStudent, ans.QuestionId)
	return ans
}
