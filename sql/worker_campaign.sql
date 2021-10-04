create table if not exists worker_campaign
(
    worker_id uuid,
    campaign_id uuid,
    foreign key(worker_id) references workers(id),
    foreign key(campaign_id) references campaigns(id),
    primary key(worker_id, campaign_id)
)