package service

import (
	httpError "backend/Error"
	"backend/database"
	"backend/helpers"
	"backend/models"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func CreateUser(user *models.User, ctx context.Context, provider string) error {
	//set Provider
	user.Provider = provider
	//set UserType
	if user.User_type == nil {
		var userType = "USER"
		user.User_type = &userType
	}
	// Validate input only for Local provider
	if provider == "Local" {
		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			return httpError.ErrInvalidInput
		}
	}
	// check email exist
	if exists, err := checkExists(ctx, "email", user.Email); err != nil {
		return httpError.ErrInternal
	} else if exists {
		return httpError.New(httpError.ErrBadRequest.StatusCode, "Email already exists")
	}
	// check Phone number exist
	if exists, err := checkExists(ctx, "phone", user.Phone); err != nil {
		return httpError.ErrInternal
	} else if exists {
		return httpError.New(httpError.ErrBadRequest.StatusCode, "Phone number already exists")
	}
	//set password
	if user.Password == nil {
		if provider == "Local" {
			return httpError.New(httpError.ErrBadRequest.StatusCode, "Password is required for local accounts")
		}
		randomPassword := helpers.GenerateRandomPassword(12)
		user.Password = &randomPassword
	}
	// hash password
	password := helpers.HashPassword(*user.Password)
	user.Password = &password
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	//set userId
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	//generater token
	token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, *user.User_type, *&user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken
	//Insert to database
	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return httpError.ErrInternal
	}
	return nil
}

func LoginUser(user models.User, ctx context.Context, userCollection *mongo.Collection) (models.User, string, string, error) {
	foundUser, err := FindUserByEmail(*user.Email, ctx, userCollection)
	if err != nil {
		return models.User{}, "", "", httpError.ErrLoginFailed
	}
	if foundUser.Provider != "Local" {
		token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, *&foundUser.User_id)
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		return *foundUser, token, refreshToken, nil
	}
	passwordIsValid, _ := helpers.VerifyPassword(*user.Password, *foundUser.Password)
	if !passwordIsValid {
		return models.User{}, "", "", httpError.ErrLoginFailed
	}
	token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, *&foundUser.User_id)
	helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)
	return *foundUser, token, refreshToken, nil
}

func FindUserByEmail(email string, ctx context.Context, userCollection *mongo.Collection) (*models.User, error) {
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}
	return &user, nil
}

func checkExists(ctx context.Context, field string, value *string) (bool, error) {
	if value == nil {
		return false, nil
	}
	count, err := userCollection.CountDocuments(ctx, bson.M{field: value})
	if err != nil {
		return false, fmt.Errorf("error checking %s: %w", field, err)
	}
	return count > 0, nil
}

func FindUserByID(ctx context.Context, uid string) (*models.User, error) {
	redisKey := fmt.Sprintf("user:%s", uid)
	val, err := database.Rdb.Get(ctx, redisKey).Result()
	if err == nil {
		var cachedUser models.User
		if err := jsoniter.Unmarshal([]byte(val), &cachedUser); err == nil {
			return &cachedUser, nil
		}
	}
	var userInfo models.User
	filter := bson.M{"user_id": uid}
	err = userCollection.FindOne(ctx, filter).Decode(&userInfo)
	if err != nil {
		return nil, err
	}
	jsonBytes, _ := jsoniter.Marshal(userInfo)
	database.Rdb.Set(ctx, redisKey, jsonBytes, 2*time.Hour)

	return &userInfo, nil
}

func FindUsersByIDs(ctx context.Context, ids []string) (map[string]string, error) {
	filter := bson.M{"user_id": bson.M{"$in": ids}}

	cursor, err := userCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("mongo find error: %w", err)
	}
	defer cursor.Close(ctx)

	result := make(map[string]string)
	for cursor.Next(ctx) {
		var u models.User
		if err := cursor.Decode(&u); err != nil {
			log.Println("decode failed:", err)
			continue
		}
		first := ""
		last := ""
		if u.First_name != nil {
			first = *u.First_name
		}
		if u.Last_name != nil {
			last = *u.Last_name
		}
		result[u.User_id] = first + " " + last
	}
	return result, nil
}

func UpdateInfo(ctx context.Context, searchField string, searchValue string, updateField string, newInfo string) error {
	exists, err := checkExists(ctx, searchField, &searchValue)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("document with %s=%s not found", searchField, searchValue)
	}

	filter := bson.M{searchField: searchValue}
	update := bson.M{
		"$set": bson.M{
			updateField: newInfo,
		},
	}

	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// func ListUsers(page int, recordPerPage int) ([]models.User, int, error)
