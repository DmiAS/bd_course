create table if not exists clients
(
    foreign key(id) references ids(id) on delete cascade
) inherits(users);