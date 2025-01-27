CREATE DATABASE musicstore;
USE musicstore;

CREATE TABLE artists (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    country VARCHAR(100),
    formed_year INT
);

CREATE TABLE albums (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    release_year INT,
    price DECIMAL(10,2),
    artist_id INT,
    FOREIGN KEY (artist_id) REFERENCES artists(id)
);

CREATE TABLE songs (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    album_id INT,
    duration VARCHAR(8),
    year INT,
    FOREIGN KEY (album_id) REFERENCES albums(id)
);