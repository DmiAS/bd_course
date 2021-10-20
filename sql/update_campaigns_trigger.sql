drop table if exists campaign_changes;

create table campaign_changes
(
    camp_id       uuid,
    prev_id       uuid,
    new_id        uuid,
    attached_type text,
    change_date   date,
    change_time   time
);

create or replace function log_campaign_change() returns trigger as
$$
declare
    prev_id       uuid;
    new_id        uuid;
    attached_type text = '';
begin
    if new.targetolog_id != old.targetolog_id then
        prev_id = old.targetolog_id;
        new_id = new.targetolog_id;
        attached_type = 'change worker';
    end if;
    if new.thread_id != old.thread_id then
        prev_id = old.thread_id;
        new_id = new.thread_id;
        attached_type = 'change thread';
    end if;
    if attached_type != '' then
        insert into campaign_changes
        values (new.id, prev_id, new_id, attached_type, current_date, current_time);
    end if;
    return new;
end;
$$ language plpgsql;

drop trigger if exists log_campaign_change on campaigns;

create trigger log_campaign_change
    after update
    on campaigns
    for each row
execute function log_campaign_change();