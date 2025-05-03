package helpers

import (
	"backend/models"
	"fmt"
	"math/rand/v2"
	"strings"
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

func CheckResult(result models.Result) string {
	if result.A == 4 {
		return "Congratulations, you won the game!"
	}
	return fmt.Sprintf("%dA%dB", result.A, result.B)
}
func CheckAnswer(answer, guess string) models.Result {
	result := models.Result{}
	answerSlice := strings.Split(answer, "")
	guessSlice := strings.Split(guess, "")
	usedAnswer := make([]bool, len(answerSlice))
	usedGuess := make([]bool, len(guessSlice))

	for i := 0; i < len(answerSlice); i++ {
		if answerSlice[i] == guessSlice[i] {
			result.A++
			usedAnswer[i] = true
			usedGuess[i] = true
		}
	}

	for i := 0; i < len(answerSlice); i++ {
		if usedAnswer[i] {
			continue
		}
		for j := 0; j < len(guessSlice); j++ {
			if !usedGuess[j] && answerSlice[i] == guessSlice[j] && i != j {
				result.B++
				usedGuess[j] = true
				break
			}
		}
	}

	return result
}
