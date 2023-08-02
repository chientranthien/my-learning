create schema if not exists be_db;

use be_db;

create table if not exists order_tab (
    id bigint(20) not null default 0,
    region varchar(20) not null default '',
    user_id bigint(20) not null default 0,
    status int(11) not null default 0,
    extra_info blob ,
    ctime bigint(20) not null default 0,
    mtime bigint(20) not null default 0
);