package controllers

import (
	"backend/helpers"
	"backend/models"
	"backend/service"
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GoogleLogin(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}
func GoogleCallback(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	var u models.User
	var foundUser models.User
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// Check Email Existing
	err = userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()
	//First Login
	if err != nil {
		//Create New User
		u.Email = &user.Email
		randomPasssword := GenerateRandomPassword(12)
		password := HashPassword(randomPasssword)
		u.Password = &password
		u.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		u.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		u.ID = primitive.NewObjectID()
		u.User_id = u.ID.Hex()
		userType := "USER"
		u.User_type = &userType
		if user.FirstName != "" {
			u.First_name = &user.FirstName
		} else {
			u.First_name = &user.Name
		}
		emptyLastName := ""
		u.Last_name = &emptyLastName
		token, refreshToken, _ := helpers.GenerateAllTokens(*u.Email, *u.First_name, *u.Last_name, *u.User_type, u.User_id)
		u.Token = &token
		u.Refresh_token = &refreshToken
		_, insertErr := userCollection.InsertOne(ctx, u)
		helpers.UpdateAllTokens(token, refreshToken, u.User_id)
		if insertErr != nil {
			apiErr := service.NewError(service.ErrInternalFailure, insertErr)
			panic(apiErr)
		}
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/auth/callback?token="+token)
		return
	}
	token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)
	helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/auth/callback?token="+token)
	return

}

func GenerateRandomPassword(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		log.Panic(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
