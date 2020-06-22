-- INSERT INTO users (id, first_name, last_name) VALUES (1, 'gopher db 1', 'gopher db 2');
-- INSERT INTO orders (id, user_id) VALUES (1, 1);
-- INSERT INTO orders (id, user_id) VALUES (2, 1);
CREATE TABLE users (
	id INTEGER NOT NULL
	,first_name VARCHAR(60) NOT NULL
	,last_name VARCHAR(60) NOT NULL
	,hashed_password VARCHAR(140)
	);

ALTER TABLE users ADD CONSTRAINT user_pkey PRIMARY KEY (id);

CREATE TABLE orders (
	id INTEGER NOT NULL
	,user_id INTEGER NOT NULL
	);

ALTER TABLE orders ADD CONSTRAINT order_pkey PRIMARY KEY (id);

ALTER TABLE orders ADD CONSTRAINT user_orders_fkey FOREIGN KEY (user_id) REFERENCES users (id);

CREATE TABLE chats (
	id INTEGER NOT NULL
	,uuid VARCHAR(55) NOT NULL
	,total_unread_messages INTEGER DEFAULT 0
	);

ALTER TABLE chats ADD CONSTRAINT chat_pkey PRIMARY KEY (id);
