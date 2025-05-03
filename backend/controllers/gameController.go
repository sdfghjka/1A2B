package controllers

import (
	httpError "backend/Error"
	"backend/database"
	"backend/helpers"
	"backend/models"
	"backend/service"
	"fmt"
	"net/http"
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
			panic(httpError.New(httpError.ErrInternal.StatusCode, "JSON encode falied"))
		}
		err = database.Rdb.Set(database.Ctx, userId, val, 5*time.Minute).Err()
		if err != nil {

			panic(httpError.New(httpError.ErrInternal.StatusCode, "Redis 寫入失敗"))
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

			panic(httpError.New(httpError.ErrInternal.StatusCode, "Can't accept your answer"))
		}
		userId := c.GetString("uid")
		data, err := database.Rdb.Get(database.Ctx, userId).Result()
		if err != nil {

			panic(httpError.New(httpError.ErrInternal.StatusCode, "Get data from Redis failed"))
		}
		var user models.GameUser
		err = jsoniter.Unmarshal([]byte(data), &user)
		user.IncreaseCount()
		result := helpers.CheckAnswer(user.Answer, body.Guess)
		if result.A == 4 {
			var u models.OverUser
			u.ID = userId
			database.Rdb.Del(database.Ctx, userId)
			duration := time.Since(user.StartTime).Seconds()
			u.Time = duration
			u.Count = user.Count
			c.JSON(http.StatusOK, gin.H{
				"result":    "Congratulations, you won the game!",
				"guess":     body.Guess,
				"count":     user.Count,
				"duration":  fmt.Sprintf("%.2f seconds", duration),
				"startTime": user.StartTime.Format(time.RFC3339),
			})

			us := c.MustGet("userService").(service.UserService)
			err = us.InsertUser(u)
			if err != nil {

				panic(httpError.New(httpError.ErrInternal.StatusCode, "Failed to insert user"))
			}
			return
		}

		val, _ := jsoniter.Marshal(user)
		database.Rdb.Set(database.Ctx, userId, val, 5*time.Minute)
		c.JSON(http.StatusOK, gin.H{"result": helpers.CheckResult(result), "guess": body.Guess})
	}
}

func GetRank() gin.HandlerFunc {
	return func(c *gin.Context) {

		us := c.MustGet("userService").(service.UserService)
		user, err := us.OrderByCount()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, user)
	}
}
