CREATE TABLE customers
(
    id serial not null unique,
    name varchar(255) not null,
    surname varchar(255) not null,
    age int not null,
    balance int not null,

    phone varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE storage(
    storage int
)