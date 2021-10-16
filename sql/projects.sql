create table if not exists projects
(
    id uuid primary key,
    client_id uuid,
    name varchar(128) not null unique ,
    foreign key(client_id) references users(id)
);