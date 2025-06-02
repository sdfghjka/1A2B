CREATE TABLE user_friends (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_uid CHAR(24) NOT NULL,
  friend_uid CHAR(24) NOT NULL,
  status ENUM('pending', 'accepted', 'blocked') DEFAULT 'pending',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_friend_user FOREIGN KEY (user_uid) REFERENCES game_users(uid),
  CONSTRAINT fk_friend_target FOREIGN KEY (friend_uid) REFERENCES game_users(uid),
  
  CONSTRAINT uq_friend_pair UNIQUE (user_uid, friend_uid),
  CONSTRAINT chk_not_self_friend CHECK (user_uid != friend_uid)
);
