package chapter08

import "strings"

func maxScore(s string) (score int) {
	for i, _ := range s {
		l := s[0 : i+1]
		r := s[i+1:]
		result := strings.Count(l, "0") + strings.Count(r, "1")
		if result > score {
			score = result
		}
	}

	return
}
