create table if not exists workers
(
    grade    varchar(16),
    position varchar(32)
) inherits (users);