CREATE TABLE mumbles (
	mumble_id INTEGER PRIMARY KEY,
	user_name TEXT NOT NULL,
	mumble TEXT NOT NULL,
	likes INTEGER NOT NULL,
	mumbletime INTEGER NOT NULL
);

insert into mumbles (mumble_id, user_name, mumble, likes, mumbletime) values
(1, 'alice', 'Hello, world!', 5, 1625155200),
(2, 'bob', 'Just had a great lunch.', 3, 1625241600),
(3, 'charlie', 'Looking forward to the weekend!', 8, 1625328000);