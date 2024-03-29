-- grant privileges for administrator
grant ALL
    on auths, campaign_stats, campaigns, projects, threads, users, workers, campaign_changes, updated_users
    to administrator;

-- grant privileges for client
grant select, update
    on users, auths
    to client;

grant insert
    on updated_users
    to client;

grant select
    on projects, threads, campaigns, campaign_stats
    to client;

-- grant privileges for worker
grant select
    on workers
    to worker;

grant insert
    on campaign_changes, updated_users
    to worker;

grant select, update
    on users
    to worker;

grant select, insert, update, delete
    on threads
    to worker;

grant select, update
    on campaigns
    to worker;

grant select, update
    on auths
    to worker;

grant select
    on projects, campaign_stats
    to worker;