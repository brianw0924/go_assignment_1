package main

import (
	. "github.com/brianw0924/go_assignment_1/classroom"
)

func main() {
	studentList := []*Student{}

	for _, name := range []string{"A", "B", "C", "D", "E"} {
		studentList = append(studentList, NewStudent(name))
	}

	teacher := NewTeacher()

	for _, s := range studentList {
		teacher.Register(s)
		go s.Start(teacher.QuestionChan, teacher.AnswerChan)
	}

	go teacher.Start()

	ch := make(chan struct{})
	<-ch
}
