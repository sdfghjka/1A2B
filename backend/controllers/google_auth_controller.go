package controllers

import (
	httpError "backend/Error"
	"backend/helpers"
	"backend/models"
	"backend/service"
	"context"
	"errors"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"go.mongodb.org/mongo-driver/mongo"
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
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		panic(httpError.ErrUnauthorized)
	}
	// Check Email Existing
	foundUser, err := service.FindUserByEmail(user.Email, ctx, userCollection)
	defer cancel()
	//First Login
	if err != nil {
		//Create New User
		if errors.Is(err, mongo.ErrNoDocuments) {
			u.Email = &user.Email
			var userType = "USER"
			u.User_type = &userType
			if user.FirstName != "" {
				u.First_name = &user.FirstName
			} else {
				u.First_name = &user.Name
			}
			emptyLastName := ""
			u.Last_name = &emptyLastName
			err := service.CreateUser(&u, ctx, "Google")
			if err != nil {
				panic(httpError.ErrInternal)
			}
			token, refreshToken, _ := helpers.GenerateAllTokens(*u.Email, *u.First_name, *u.Last_name, *u.User_type, u.User_id)
			u.Token = &token
			u.Refresh_token = &refreshToken
			_, insertErr := userCollection.InsertOne(ctx, u)
			helpers.UpdateAllTokens(token, refreshToken, u.User_id)
			if insertErr != nil {
				panic(httpError.ErrInternal)
			}
			c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/auth/callback?token="+token)
			return
		} else {
			panic(httpError.ErrInternal)
		}

	}
	token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, foundUser.User_id)
	helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/auth/callback?token="+token)
	return

}
