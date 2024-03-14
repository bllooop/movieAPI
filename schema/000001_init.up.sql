CREATE TABLE userlist
(
    id serial not null unique,
    username varchar(255) not null unique,
    password varchar(255) not null,
    role varchar(5) not null
);

CREATE TABLE movieitem
(
    id serial not null unique,
    title varchar(150) not null,
    description varchar(1000) not null,
    rating int not null,
    date varchar(255) not null,
    actorname text ARRAY
);

CREATE TABLE actoritem
(
    id serial not null unique,
    name varchar(100) not null,
    gender varchar(10) not null,
    date varchar(50) not null
);


CREATE TABLE movielist
(
    id serial not null unique,
    title varchar(150) not null,
    rating int not null,
    date varchar(255) not null
);

CREATE TABLE actorlist
(
    id serial not null unique,
    name varchar(100) not null,
    gender varchar(10) not null,
    date varchar(50) not null
);

CREATE TABLE movlistitem
(
    id serial not null unique,
    mov_list_id int references movieitem(id) on delete cascade not null,
    mov_item_id int references movieitem(id) on delete cascade not null,
    actor_list_id int references actorlist(id) on delete cascade not null,
    actor_item_id int references actoritem(id) on delete cascade not null
);