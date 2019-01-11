package score // import "moul.io/patfind/score"

type Sequence string

func New(sequence string) Sequence {
	return Sequence(sequence)
}

func (n Sequence) String() string {
	return string(n)
}

func (n Sequence) Score() (average float64, comments []string) {
	comments = []string{}
	var totalScores float64

	for _, handler := range EnabledHandlers {
		score, comment := handler(string(n))
		totalScores += score
		if comment != "" {
			comments = append(comments, comment)
		}
	}

	if totalScores != 0 {
		average = totalScores / float64(len(EnabledHandlers))
	}
	return average, comments
}
