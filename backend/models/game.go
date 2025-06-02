package models

import "database/sql"

type GameMatch struct {
	Player1UID string         `db:"player1_uid"`
	Player2UID string         `db:"player2_uid"`
	WinnerUID  sql.NullString `db:"winner_uid"`
}
