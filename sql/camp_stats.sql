create table if not exists camp_stats(
    camp_id uuid primary key,
    date date, -- дата сбора статистики
    spent float,
    impressions int, -- количество показов
    conversion int,
    subs int[],
    unsubs int[],
    sales int,
    foreign key(camp_id) references campaigns(id)
)