create table if not exists users
(
    id uuid primary key,
    first_name varchar(64),
    last_name varchar(64),
    vk_link text,
    tg_link text,
    role varchar(16) CHECK ( role in ('admin', 'worker', 'client')),
    created bigint not null,
    foreign key (id) references auths(user_id) on delete cascade
)