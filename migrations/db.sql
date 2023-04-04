create table courses (
  id serial PRIMARY KEY,
  title varchar(256) NOT NULL,
  price int NOT NULL,
  banner varchar(600) NOT NULL
)