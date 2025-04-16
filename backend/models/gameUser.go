package models

import "time"

type GameUser struct {
	ID        string    `json:"id"`
	StartTime time.Time `json:"start_time"`
	Count     int       `json:"count"`
	Answer    string    `json:"answer"`
}

type Result struct {
	A int
	B int
}

func NewUser(id, answer string) *GameUser {
	return &GameUser{
		ID:        id,
		StartTime: time.Now(),
		Count:     0,
		Answer:    answer,
	}
}

func (user *GameUser) IncreaseCount() {
	user.Count++
}
