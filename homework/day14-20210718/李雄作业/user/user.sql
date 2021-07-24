create database if not exists user default charset utf8mb4;

create table if not exists user (id bigint primary key auto_increment,name varchar(32) not null default '',sex boolean not null default true,addr varchar(128) not null default '',created_at datetime not null,updated_at datetime not null,deleted_at datetime) engine=innodb default charset utf8mb4;

alter table user add column remark varchar(128) not null default '';

-- 增加密码字段
alter table user add column  password varchar(1024) not null default '';
-- 只增不改，不减



update user set password='$2a$10$pbaRODVnyr2trM4vVZH3.uPBBjb/BxByc71s0DmpAgFNYWS1YeVq6';




-- 数据库修改密码

ALTER USER 'root'@'%' IDENTIFIED BY 'adl1314520';

ALTER USER 'root'@'localhost' IDENTIFIED BY 'adl1314520';

