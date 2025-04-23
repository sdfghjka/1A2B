CREATE TABLE game_users (
  uid CHAR(24) NOT NULL PRIMARY KEY,      
  time DATETIME NOT NULL,            
  count INT DEFAULT 0                                 
);