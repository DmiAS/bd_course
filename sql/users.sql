create table if not exists users
(
    id         uuid primary key,
    first_name varchar(128),
    last_name  varchar(128),
    vk_link    text,
    tg_link    text
)