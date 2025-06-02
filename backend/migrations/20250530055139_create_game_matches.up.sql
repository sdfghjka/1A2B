CREATE TABLE game_matches (
  id INT AUTO_INCREMENT PRIMARY KEY,
  player1_uid CHAR(24) NOT NULL,
  player2_uid CHAR(24) NOT NULL,
  winner_uid CHAR(24),
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_match_player1 FOREIGN KEY (player1_uid) REFERENCES game_users(uid),
  CONSTRAINT fk_match_player2 FOREIGN KEY (player2_uid) REFERENCES game_users(uid),
  CONSTRAINT fk_match_winner FOREIGN KEY (winner_uid) REFERENCES game_users(uid)
);
