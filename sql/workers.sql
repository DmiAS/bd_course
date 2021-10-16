create table if not exists workers
(
    id uuid primary key,
    grade    varchar(16),
    position varchar(32),
    foreign key (id) references auths(user_id) on delete cascade
)