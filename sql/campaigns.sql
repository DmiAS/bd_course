create table if not exists campaigns
(
    id         uuid primary key,
    thread_id  uuid,
    cabinet_id int not null,
    client_id  int not null,
    vk_camp_id int not null,
    name       varchar(256)
)