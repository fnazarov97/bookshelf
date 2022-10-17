DROP DATABASE IF EXISTS bookshelf;
CREATE DATABASE IF NOT EXISTS bookshelf;

CREATE TABLE IF NOT EXISTS users(
    "id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
    "name" VARCHAR(30) NOT NULL,
    "key" VARCHAR(50) NOT NULL,
    "secret" VARCHAR(50) NOT NULL,
     "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users(
    name, key, secret)
VALUES('Farhod','123', 'My_secret'),
      ('Doniyor','321', 'his_secret'),
      ('Ahmad','453', 'secret');

CREATE TABLE IF NOT EXISTS books(
    "id" SERIAL NOT NULL UNIQUE PRIMARY KEY,
    "isbn" text NOT NULL,
    "title" varchar(45) NOT NULL,
    "author" varchar(45) NOT NULL,
    "published" INT NOT NULL,
    "pages" INT NOT NULL,
    "status" INT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT NULL,
    "deleted_at" TIMESTAMP DEFAULT NULL
);


INSERT INTO books(
    isbn, title, author, published, pages, status)
VALUES('s412a5', 'Bahor qaytib', 'Oybek', '2000', '200', 1),
      ('s412a5', 'Kapitan Grand bolalari', 'Jyul Vern', '1990', '410', 0),
      ('s412a5', 'Bolalik', 'Oybek', '2000', '300', 0);
