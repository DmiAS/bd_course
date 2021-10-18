create table if not exists campaign_stats
(
    camp_id     uuid primary key,
    date        date,  -- дата сбора статистики
    spent       float,
    impressions int,   -- количество показов
    conversion  int,   -- количество переходов
    subs        int[], -- id подписчиков в вк
    unsubs      int[], -- id отписавшизся в вк
    sales       int,   -- количество продаж
    foreign key (camp_id) references campaigns (id) on delete cascade
)