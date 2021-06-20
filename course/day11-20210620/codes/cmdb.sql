create database if not exists `cmdb` default charset utf8mb4;

use `cmdb`;

create table if not exists `asset`(
    `id` bigint primary key auto_increment comment "主键",
    `ip` varchar(128) not null default "" comment "资产IP",
    `addr` varchar(128) not null default "" comment "资产位置",
    `created_at` datetime not null,
    `updated_at` datetime not null,
    `deleted_at` datetime
)engine=innodb default charset utf8mb4;

create table if not exists `module` (
    `id` bigint primary key auto_increment,
    `name` varchar(128) not null default "",
    `created_at` datetime not null,
    `updated_at` datetime not null,
    `deleted_at` datetime
) engine=innodb default charset utf8mb4;

create table if not exists `role` (
    `id` bigint primary key auto_increment,
    `name` varchar(128) not null default "",
    `created_at` datetime not null,
    `updated_at` datetime not null,
    `deleted_at` datetime
) engine=innodb default charset utf8mb4;

create table if not exists `rel_role_module`(
    `role_id` bigint not null default 0,
    `module_id` bigint not null default 0,
    `created_at` datetime not null,
    `updated_at` datetime not null,
    `deleted_at` datetime
) engine=innodb default charset utf8mb4;

create table if not exists `user` (
    `id` bigint primary key auto_increment,
    `name` varchar(64) not null default "",
    `password` varchar(1024) not null default "",
    `birthday` date not null,
    `telephone` varchar(64) not null default "",
    `email` varchar(64) not null default "",
    `addr` varchar(128) not null default "",
    `status` tinyint not null default 0,
    `role_id` bigint not null default 0,
    `department_id` bigint not null default 0,
    `created_at` datetime not null,
    `updated_at` datetime not null,
    `deleted_at` datetime,
    index `idx_name`(`name`)
) engine=innodb default charset utf8mb4;

alter table `user` add column `description` text not null default "";


create table if not exists `department` (
    `id` bigint primary key auto_increment,
    `name` varchar(64) not null default "",
    `created_at` datetime not null,
    `updated_at` datetime not null,
    `deleted_at` datetime
) engine=innodb default charset utf8mb4;
