package service

import (
	"backend/models"
	"database/sql"
	"fmt"
)

type UserService interface {
	GetAllUsers() ([]models.OverUser, error)
	InsertUser(M models.OverUser) error
	OrderByCount() ([]models.OverUser, error)
}

type userService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return &userService{DB: db}
}

func (s *userService) GetAllUsers() ([]models.OverUser, error) {
	rows, err := s.DB.Query("SELECT id, time, count FROM game_users")
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var users []models.OverUser
	for rows.Next() {
		var u models.OverUser
		err := rows.Scan(&u.ID, &u.Time, &u.Count)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		users = append(users, u)
	}
	return users, nil
}

func (s *userService) InsertUser(M models.OverUser) error {
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

func (s *userService) OrderByCount() ([]models.OverUser, error) {
	rows, err := s.DB.Query("SELECT * FROM game_users ORDER BY count ASC, time ASC;")
	if err != nil {
		return nil, fmt.Errorf("query failed: %w from OrderByCount", err)
	}
	defer rows.Close()
	var users []models.OverUser
	for rows.Next() {
		var u models.OverUser
		err := rows.Scan(&u.ID, &u.Time, &u.Count)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		users = append(users, u)

	}
	return users, nil
}
