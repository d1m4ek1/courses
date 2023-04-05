create table courses (
  id serial PRIMARY KEY,
  title varchar(256) NOT NULL,
  price int NOT NULL,
  banner varchar(600) NOT NULL,
  add_date text NOT NULL,
  user_id varchar(128) NOT NULL
)

create table cart (
  id serial PRIMARY KEY,
  course_id int NOT NULL,
  count int NOT NULL,
  user_id varchar(128) NOT NULL 
);

create table orders (
  id serial PRIMARY KEY,
  cart_items_id int[] NOT NULL,
  user_id varchar(128) NOT NULL
)