create table if not exists campaign_stats
(
    camp_id     uuid,
    date        date,  -- дата сбора статистики
    spent       float,
    impressions int,   -- количество показов
    conversion  int,   -- количество переходов
    subs        bigint[], -- id подписчиков в вк
    unsubs      bigint[], -- id отписавшизся в вк
    sales       int,   -- количество продаж
    foreign key (camp_id) references campaigns (id) on delete cascade,
    primary key(camp_id, date)
)