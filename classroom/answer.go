package classroom

type Answer struct {
	QuestionId  int
	fromStudent string
	val         int
}

func NewAnswer(fromStudent string, val int, questionId int) *Answer {
	return &Answer{
		QuestionId:  questionId,
		fromStudent: fromStudent,
		val:         val,
	}
}
