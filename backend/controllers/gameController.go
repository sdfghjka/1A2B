package controllers

import (
	"backend/database"
	"backend/helpers"
	"backend/models"
	"backend/service"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func GenerateAnswer() gin.HandlerFunc {
	return func(c *gin.Context) {
		answer := helpers.GenerateAnswer()
		userId := c.GetString("uid")
		user := models.NewUser(userId, answer)
		val, err := jsoniter.Marshal(user)
		if err != nil {
			panic(service.NewError(service.ErrInternalFailure, fmt.Errorf("JSON encode falied")))
		}
		err = database.Rdb.Set(database.Ctx, userId, val, 5*time.Minute).Err()
		if err != nil {
			apiErr := service.NewError(service.ErrInternalFailure, fmt.Errorf("Redis 存資料失敗"))
			panic(apiErr)
		}
		c.JSON(http.StatusOK, gin.H{"userId": userId})
	}
}

func Guess() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Guess string `json:"guess"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			apiErr := service.NewError(service.ErrInternalFailure, fmt.Errorf("Can't accept your answer"))
			panic(apiErr)
		}
		userId := c.GetString("uid")
		data, err := database.Rdb.Get(database.Ctx, userId).Result()
		if err != nil {
			apiErr := service.NewError(service.ErrInternalFailure, fmt.Errorf("Failed to get game data"))
			panic(apiErr)
		}
		var user models.GameUser
		err = jsoniter.Unmarshal([]byte(data), &user)
		user.IncreaseCount()
		result := CheckAnswer(user.Answer, body.Guess)
		if result.A == 4 {
			database.Rdb.Del(database.Ctx, userId)
			duration := time.Since(user.StartTime).Seconds()

			c.JSON(http.StatusOK, gin.H{
				"result":    "Congratulations, you won the game!",
				"guess":     body.Guess,
				"count":     user.Count,
				"duration":  fmt.Sprintf("%.2f seconds", duration),
				"startTime": user.StartTime.Format(time.RFC3339),
			})
			return
		}

		val, _ := jsoniter.Marshal(user)
		database.Rdb.Set(database.Ctx, userId, val, 5*time.Minute)
		c.JSON(http.StatusOK, gin.H{"result": CheckResult(result), "guess": body.Guess})
	}
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
