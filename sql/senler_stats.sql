-- таблица заполняется в течение дня всеми заявками, даже если кампания не была добавлена
-- в систему, в 00 таблица за предудыщий день очищается и заполняет оставшуюсю информацию
-- в таблице camp_stats
create table if not exists senler_stats(
    id uuid primary key,
    camp_id int,
    sub_type varchar(5) check (sub_type in ('sub', 'unsub')),
    sub_id int -- vk_id подписчика
)