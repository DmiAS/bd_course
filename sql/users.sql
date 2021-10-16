create table if not exists users
(
    id uuid primary key,
    first_name varchar(64),
    last_name varchar(64),
    vk_link text,
    tg_link text,
    foreign key (id) references auths(user_id) on delete cascade
)