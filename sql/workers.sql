create table if not exists workers
(
    id         uuid,
    first_name varchar(128),
    last_name  varchar(128),
    grade      varchar(16),
    position   varchar(32),
    vk_link    text,
    tg_link    text
);