package helpers

import (
	"backend/models"
	"fmt"
)

func GenerateAllAnswer() []string {
	var res []string
	for i := 0; i < 9999; i++ {
		s := fmt.Sprintf("%04d", i)
		if isUniqueDigits(s) {
			res = append(res, s)
		}
	}
	return res
}

func isUniqueDigits(s string) bool {
	seen := make(map[rune]bool)
	for _, ch := range s {
		if seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

func FilterPossibleAnswers(candidates []string, guess string, result models.Result) []string {
	var filtered []string
	for _, ans := range candidates {
		if CheckAnswer(ans, guess) == result {
			filtered = append(filtered, ans)
		}
	}
	return filtered
}
