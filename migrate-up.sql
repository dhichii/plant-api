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
VALUES ('super', 'super@company.com', '$2a$04$w4mbUkxHup/8TqTZDVlmFO9IviXHcHFPu/3KKp9UKDkCaTenOKS0O', 'super', NOW(), NOW());

CREATE TABLE plants(
  id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(255),
  botanical_name VARCHAR(255),
  type VARCHAR(20),
  description LONGTEXT,
  difficulty VARCHAR(6),
  watering_time VARCHAR(15),
  how_to_grow LONGTEXT,
  soil LONGTEXT,
  created_at DATETIME(3),
  updated_at DATETIME(3),
  deleted_at DATETIME(3)
);

CREATE TABLE natives(
  id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(120),
  created_at DATETIME(3),
  updated_at DATETIME(3),
  deleted_at DATETIME(3)
);

CREATE TABLE plant_natives(
  plant_id BIGINT UNSIGNED NOT NULL,
  native_id BIGINT UNSIGNED NOT NULL,
  PRIMARY KEY(plant_id, native_id),
  CONSTRAINT fk_1_plant_natives_plant FOREIGN KEY (plant_id) REFERENCES plants(id),
  CONSTRAINT fk_2_plant_natives_native FOREIGN KEY (native_id) REFERENCES natives(id),
  INDEX (plant_id)
);