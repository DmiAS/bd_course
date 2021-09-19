create table if not exists passwords (
    login varchar(32),
    password varchar(256),
    salt varchar(32)
);