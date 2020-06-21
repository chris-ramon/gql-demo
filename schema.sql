-- INSERT INTO users (id, first_name, last_name) VALUES (1, 'gopher db 1', 'gopher db 2');
-- INSERT INTO orders (id, user_id) VALUES (1, 1);
-- INSERT INTO orders (id, user_id) VALUES (2, 1);

CREATE TABLE users (
  id 					integer NOT NULL,
  first_name 	varchar(60) NOT NULL,
  last_name 	varchar(60) NOT NULL,
  hashed_password varchar(140)
);

ALTER TABLE users ADD CONSTRAINT user_pkey PRIMARY KEY (id);

CREATE TABLE orders (
  id 					integer NOT NULL,
  user_id integer NOT NULL
);

ALTER TABLE orders ADD CONSTRAINT order_pkey PRIMARY KEY (id);
ALTER TABLE orders ADD CONSTRAINT user_orders_fkey FOREIGN KEY (user_id) REFERENCES users(id);

CREATE TABLE chats (
  id    integer NOT NULL,
  uuid 	varchar(55) NOT NULL
);

ALTER TABLE chats ADD CONSTRAINT chat_pkey PRIMARY KEY (id);
