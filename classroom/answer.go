package classroom

type Answer struct {
	fromStudent string
	val         int
}

func NewAnswer(fromStudent string, val int) *Answer {
	return &Answer{
		fromStudent,
		val,
	}
}
