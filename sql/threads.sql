create table if not exists threads(
    id uuid primary key,
    project_id uuid,
    name varchar(128)
);