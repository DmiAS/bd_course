create table if not exists threads
(
    id         uuid primary key,
    project_id uuid,
    name       varchar(128) not null,
    created    bigint       not null,
    foreign key (project_id) references projects (id) on delete cascade,
    unique (project_id, name)
);