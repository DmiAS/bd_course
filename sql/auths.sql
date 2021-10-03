create table if not exists auths (
    login varchar(32) unique,
    password varchar(256),
    salt varchar(32),
    user_id uuid,
    foreign key(user_id) references ids(id) on delete cascade
);