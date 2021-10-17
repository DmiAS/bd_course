create table if not exists workers
(
    user_id uuid primary key,
    grade    varchar(16),
    position varchar(32),
    foreign key (user_id) references users(id) on delete cascade
)