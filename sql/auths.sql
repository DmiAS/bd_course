create table if not exists auths (
    login varchar(32) unique not null,
    password varchar(256) not null,
    salt varchar(32) not null,
    role varchar(16) CHECK ( role in ('admin', 'worker', 'client')),
    user_id uuid primary key
);