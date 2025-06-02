package service

import (
	"backend/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type GameService interface {
	InsertGameRecord(record models.GameMatch) error
}

type gameService struct {
	DB *sqlx.DB
}

func NewGameService(db *sqlx.DB) GameService {
	return &gameService{DB: db}
}

func (s *gameService) InsertGameRecord(record models.GameMatch) error {
	_, err := s.DB.NamedExec("INSERT INTO game_matches (player1_uid, player2_uid, winner_uid) VALUES (:player1_uid, :player2_uid, :winner_uid)", &record)
	if err != nil {
		return fmt.Errorf("insert game record failed: %w", err)
	}
	return nil
}
