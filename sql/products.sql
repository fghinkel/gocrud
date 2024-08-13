create table if not exists products (
	id serial primary key,
	name varchar,
	description varchar,
	price decimal,
	quantity integer
);