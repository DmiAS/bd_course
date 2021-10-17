create table if not exists projects
(
    id        uuid primary key,
    client_id uuid,
    name      varchar(128) not null,
    created   bigint       not null,
    foreign key (client_id) references users (id),
    unique (client_id, name)
);