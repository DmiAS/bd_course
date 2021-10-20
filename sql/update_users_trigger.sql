drop table if exists updated_users;
create table updated_users
(
    user_id        uuid,
    changed_fields text[],
    change_date    date,
    change_time    time
);

create or replace function log_user_change() returns trigger as
$$
declare
    fields text[];
begin
    if new.first_name != old.first_name then
        fields = array_append(fields, 'first_name');
    end if;
    if new.last_name != old.last_name then
        fields = array_append(fields, 'last_name');
    end if;
    if new.vk_link != old.vk_link then
        fields = array_append(fields, 'vk_link');
    end if;
    if new.tg_link != old.tg_link then
        fields = array_append(fields, 'tg_link');
    end if;
    insert into updated_users values (new.id, fields, current_date, current_time);
    return new;
end;
$$ language plpgsql;

drop trigger if exists log_user_change on users;

create trigger log_user_change
    after update
    on users
    for each row
execute function log_user_change();