package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil
	if uid != userId {
		err = errors.New("Unauthorized to access this resource")
	}
	err = CheckUserType(c, userType)
	return err
}

func HashPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(hashPassword)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = fmt.Sprintf("Email or Password is incorrect")
		check = false
	}
	return check, msg
}

func GenerateRandomPassword(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		log.Panic(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
