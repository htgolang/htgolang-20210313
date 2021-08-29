create table osquery_agent(
    id bigint primary key auto_increment,
    uuid varchar(64) not null default '',
    hostname varchar(64) not null default '',
    addr varchar(256) not null default '',
    config longtext,
    config_version varchar(64) not null default '',
    heartbeat_time datetime not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
) engine=innodb default charset utf8mb4;


create table osquery_job (
    id bigint primary key auto_increment,
    uuid varchar(64) not null default '',
    `type` varchar(64) not null default '',
    content longtext,
    `status` int,
    result longtext,
    complated_at datetime,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
) engine=innodb default charset utf8mb4;