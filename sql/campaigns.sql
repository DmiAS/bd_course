create table if not exists campaigns
(
    id            uuid primary key,
    thread_id     uuid,
    targetolog_id uuid,
    cabinet_id    int          not null,
    client_id     int          not null,
    vk_camp_id    int          not null,
    name          varchar(256) not null,
    created       bigint not null,
    unique (thread_id, name),
    foreign key (thread_id) references threads (id),
    foreign key (targetolog_id) references workers (id)
)