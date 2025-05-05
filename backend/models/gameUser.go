package models

import "time"

type GameUser struct {
	ID        string    `json:"id"`
	StartTime time.Time `json:"start_time"`
	Count     int       `json:"count"`
	Answer    string    `json:"answer"`
}

type Result struct {
	A int `json:"a"`
	B int `json:"b"`
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

type RankedUser struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Time  float64 `json:"duration"`
	Count int     `json:"count"`
}
