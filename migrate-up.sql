CREATE TABLE users(
  id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(120),
  email VARCHAR(255) UNIQUE,
  password VARCHAR(255),
  role VARCHAR(5),
  created_at DATETIME(3),
  updated_at DATETIME(3),
  deleted_at DATETIME(3)
);

INSERT users(name, email, password, role, created_at, updated_at)
VALUES ('super', 'super@company.com', 'super123', 'super', NOW(), NOW());