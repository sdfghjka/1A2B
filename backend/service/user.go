package service

import (
	"backend/models"
	"context"
	"database/sql"
	"fmt"
)

type UserService interface {
	GetAllUsers() ([]models.RankedUser, error)
	InsertUser(M models.RankedUser) error
	GetRankedUsers(ctx context.Context) ([]models.RankedUser, error)
}

type userService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return &userService{DB: db}
}

func (s *userService) GetAllUsers() ([]models.RankedUser, error) {
	rows, err := s.DB.Query("SELECT id, time, count FROM game_users")
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var users []models.RankedUser
	for rows.Next() {
		var u models.RankedUser
		err := rows.Scan(&u.ID, &u.Time, &u.Count)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		users = append(users, u)
	}
	return users, nil
}

func (s *userService) InsertUser(M models.RankedUser) error {
	fmt.Println("Inserting user with ID:", M.ID)
	insertStmt, err := s.DB.Prepare("INSERT INTO game_users (uid, time, count) VALUES (?, ?, ?);")
	if err != nil {
		return fmt.Errorf("prepare insert failed: %w", err)
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(M.ID, M.Time, M.Count)
	if err != nil {
		return fmt.Errorf("exec insert failed: %w", err)
	}
	fmt.Println("Insert successful!")
	return nil
}

func (s *userService) GetRankedUsers(ctx context.Context) ([]models.RankedUser, error) {
	rows, err := s.DB.Query("SELECT uid, time, count FROM game_users ORDER BY count ASC, time ASC;")
	if err != nil {
		return nil, fmt.Errorf("query failed: %w from OrderByCount", err)
	}
	defer rows.Close()
	type scoreRow struct {
		ID    string
		Time  float64
		Count int
	}
	var scores []scoreRow
	var ids []string
	for rows.Next() {
		var s scoreRow
		if err := rows.Scan(&s.ID, &s.Time, &s.Count); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		scores = append(scores, s)
		ids = append(ids, s.ID)
	}
	userMap, err := FindUsersByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	var result []models.RankedUser
	for _, s := range scores {
		result = append(result, models.RankedUser{
			ID:    s.ID,
			Name:  userMap[s.ID],
			Time:  s.Time,
			Count: s.Count,
		})
	}
	return result, nil
}
