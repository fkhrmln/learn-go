-- Active: 1717342743334@@127.0.0.1@3306@go_restful_api
CREATE TABLE product (
    id VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

SELECT * FROM product;