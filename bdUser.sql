
create table users (
    user_id  serial not null PRIMARY KEY,
    username VARCHAR (50) UNIQUE,
    email VARCHAR (50) UNIQUE,
    password text,
    avatar text
);