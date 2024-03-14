CREATE TABLE userlist
(
    id serial not null unique,
    username varchar(255) not null,
    password varchar(255) not null,
    role varchar(5) not null
);
