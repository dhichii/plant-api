CREATE TABLE users(
  id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(120),
  email VARCHAR(255) UNIQUE,
  password VARCHAR(255),
  is_admin BOOLEAN,
  created_at DATETIME(3),
  updated_at DATETIME(3),
  deleted_at DATETIME(3)
);