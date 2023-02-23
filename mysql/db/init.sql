-- MySQL dump 10.13  Distrib 5.1.51, for pc-linux-gnu (i686)
--
-- Host: 127.0.0.1    Database: world
-- ------------------------------------------------------
-- Server version       5.1.51-debug-log

/*以下Databaseの内容が記載される*/

DROP SCHEMA IF EXISTS golang;
CREATE SCHEMA golang;
USE golang;

DROP TABLE IF EXISTS golang;

CREATE TABLE users
(
  id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name                VARCHAR(40),
  age                 INT(10),
  created_at          TIMESTAMP DEFAULT NULL,
  updated_at          TIMESTAMP DEFAULT NULL,
  deleted_at          TIMESTAMP DEFAULT NULL
);

INSERT INTO users (id, name, age) VALUES (1, "testKun", 21);
