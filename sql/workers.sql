create table if not exists workers
(
    user_id uuid primary key,
    grade    varchar(16),
    position varchar(32) default 'targetolog',
    foreign key (user_id) references users(id) on delete cascade
)