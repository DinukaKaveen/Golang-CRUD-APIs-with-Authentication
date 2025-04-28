CREATE TABLE authors (
  id         BIGINT       NOT NULL AUTO_INCREMENT PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  last_name  VARCHAR(255) NOT NULL,
  bio        text
);

CREATE TABLE books (
  id        BIGINT       NOT NULL AUTO_INCREMENT PRIMARY KEY,
  title     VARCHAR(255) NOT NULL,
  author_id BIGINT       NOT NULL,
  FOREIGN KEY (author_id) REFERENCES authors(id)
);