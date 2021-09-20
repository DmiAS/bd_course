create table if not exists projects
(
    id uuid primary key,
    client_id uuid,
    name varchar(128)
);