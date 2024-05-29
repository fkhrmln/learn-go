-- Active: 1716827847738@@127.0.0.1@3306@go_database
CREATE TABLE user (
    id VARCHAR(5) NOT NULL,
    name VARCHAR(50) NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

SELECT * FROM `user`;

DELETE FROM `user`;

ALTER TABLE `user`
ADD COLUMN email VARCHAR(50),
ADD COLUMN balance INTEGER DEFAULT(0),
ADD COLUMN rating DOUBLE DEFAULT(0.0),
ADD COLUMN married BOOLEAN DEFAULT FALSE,
ADD COLUMN birthdate DATE,
ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

UPDATE `user`
SET
    email = NULL,
    birthdate = NULL
WHERE
    id = "00001";

CREATE TABLE admin (
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    PRIMARY KEY (username)
) ENGINE = InnoDB;

SELECT * FROM admin;

INSERT INTO admin (username, password) VALUES ("admin", "admin");

CREATE TABLE comment (
    id INTEGER NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

SELECT * FROM comment;

DROP TABLE comment;