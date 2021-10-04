create table if not exists workers
(
    grade    varchar(16),
    position varchar(32),
    foreign key (id) references ids(id) on delete cascade,
    UNIQUE(id)
) inherits (users);