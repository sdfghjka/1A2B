package helpers

import (
	"fmt"
	"math/rand/v2"
)

func GenerateAnswer() string {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(digits), func(i, j int) {
		digits[i], digits[j] = digits[j], digits[i]
	})
	answer := ""
	for i := 0; i < 4; i++ {
		if i == 0 && digits[i] == 0 {
			return GenerateAnswer()
		}
		answer += fmt.Sprintf("%d", digits[i])
	}
	return answer
}
