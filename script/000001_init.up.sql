CREATE TABLE users
(
    id serial  not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    pswd_hash varchar(255) not null
);

CREATE TABLE todo_list
(
    id serial  not null unique,
    title varchar(255) not null,
    descript varchar(255)
);

CREATE TABLE user_list
(
    id serial  not null unique,
    user_id int references users (id) on delete cascade not null,
    list_id int references todo_list (id) on delete cascade not null
);

CREATE TABLE todo_item
(
    id serial  not null unique,
    title varchar(255) not null,
    descript varchar(255),
    done boolean not null default false
);

CREATE TABLE list_item
(
    id serial  not null unique,
    user_id int references todo_item (id) on delete cascade not null,
    list_id int references todo_list (id) on delete cascade not null
);