package main

type Number string

func New(number string) Number {
	return Number(number)
}

func (n Number) String() string {
	return string(n)
}

func (n Number) Score() (average float64, comments []string) {
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
