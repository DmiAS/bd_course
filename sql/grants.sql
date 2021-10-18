-- grant privileges for administrator
grant ALL
    on auths, campaign_stats, campaigns, projects, threads, users, workers
    to administrator;

-- grant privileges for client
grant select, update
    on users
    to client;

grant select
    on auths, projects, threads, campaigns, campaign_stats
    to client;

-- grant privileges for worker
grant select, update
    on workers
    to worker;

grant select, insert, update
    on threads
    to worker;

grant select, update
    on campaigns
    to worker;

grant select
    on auths, campaigns, campaign_stats
    to worker;