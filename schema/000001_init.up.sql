CREATE TABLE customers
(
    id serial not null unique,
    name varchar(255) not null,
    surname varchar(255) not null,
    age int not null,
    balance int not null,
<<<<<<< HEAD
=======
    phone varchar(255) not null,
>>>>>>> f48611f (first commit)
    password_hash varchar(255) not null
);