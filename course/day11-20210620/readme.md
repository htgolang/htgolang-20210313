1. 复习 10分钟
    html 超文本标记语言

    html模板
    go text/template html/template(web 开发一定要用html)
    模板语法:
        显示
        条件
        循环
        函数调用

        模板复用:
            定义 define
            引用 template
            覆盖 block

    html/template

    认证:
        跟踪用户的状态
        服务端跟踪
            session => sid => cookie
        客户端跟踪:
            编码
            加密

            对发方给客户端的数据进行签名 => 对数据签名验证
            cookie gob编码的 => 其他人拿到信息肯定可以解码 => 不能设置隐私信息

            jwt

    rpc http(restapi)

    restapi 风格 => http协议
    http => method url header => body
        method => 操作动作 增/删/改/查
            POST => 增
            DELETE => 删
            GET => 查
            PUT => 改
        url => 资源 /users/
            DELETE => ?pk/id= /users/{pk}/ /users/1/
            PUT => ?pk/id /users/{pk}/
            {
                "id" : "xxx"
            }
            GET =>
                ?pk/id = /users/{pk}/ => 查询某一个
                无 pk/id => 查询所有 ?query=&age=
                根据条件查询列表信息 []
                根据ID获取对应用户信息 一个元素

    前端/其他业务方
        rpc => http协议
               tcp协议
        从功能上没有区别
        实现上来说: tcp/udp => 数据(编码方法)
                    不同的编码方式=> 序列化/反序列化的效率 网络传输大小

    内/外:
        对内: 效率 => rpc/resetapi
        对外: 规范 => resetapi

        将内部API => 发布到 => 外网(外部API)

        resetapi =>
            功能: resetapi/rpc
            权限限制

            网关(可能用的第三方/可能自主开发) proxy
                外部resetapi => 转换为内部 rpc/resetapi
                     /url => /inner url => nginx location
                     接收客户端(http)请求 => 转码 => 发送请求(http/rpc)到服务端
                     接收服务端(http/rpc)响应 => 转码 => 给客户端(http)响应

                权限控制/限制频率/...

    rpc/restapi 是否在同一个服务(ip:port)
        自己实现:
            restapi => http
            rpc => http(gob)、tcp(gob json)

            如果说在同一个IP:port 即要发布restapi又要发布rpc
                http => method url request body => response body

            http => url => handler/handlerfunc

            rpc.HandleHttp() /rpc/

            http.Handle(/xxx/, handler())

        grpc/其他框架


    个人建议
        启动两个不同的服务 => 业务代码可以一份

        main.go => rpc/http

    个人建议：
        看一些使用go开发的工具/框架/go源码(用go语言写的)

        socket -> http http包
        rabbitmq业务逻辑 => go实现 => nsq

        => 学习别人怎么写代码

2. 作业
3. 新内容
    embed
        无html模板文件之前 => .go => 二进制程序 exe/xxxx => 部署在对应架构的环境可以运行
        用户管理 .json/.password => .go -> exe => 只部署exe+.json/.password
        html模板文件 => .go => exe => 部署exe + html
    数据库 => mariadb(mysql)
        基本知识：
            数据库实例 => mysql进程/pg_ctl进程 ip:port => 某台电脑
            数据库 => 连接数据库实例后创建库 => excel文件
            数据表 => 每个库中可以由多个表, 二维表 => sheet
            数据行 => 每行表示一个数据(数据的属性名、属性类型) => 1,2,3...
            数据列 => 每行由多个属性，每个属性都由属性名、属性类型 => A, B, C
        SQL => 结构化查询语言 => 对数据库操作 => 标准规范 => 数据库厂商
            mysql -u root -p

            DCL => 数据控制语言
                用户，权限
                create user "golang"@"%" identified by "golang@2021";
                drop user golang@'%';

                grant all privileges on db.table to golang@'%';
                revoke all privileges on db.table from golang@'%';
                show grants for golang@'%';

            DDL => 数据定义语言
                数据库操作
                    创建库
                        create database cmdb default charset utf8mb4;
                    删除库
                        drop database cmdb;
                    使用:
                        use cmdb;
                    查看库
                        show create database cmdb;
                数据表的操作
                    创建表:
                        create table tableName (
                            name type 修饰 注释 [逗号分隔, 最后一个元素后面没由逗号],
                            约束, 索引
                        ) engine=innodb default charset utf8mb4;

                        标识符: 小写英文字母 _ 数字
                        type:
                            数值
                                布尔值
                                    boolean => true/false
                                整数
                                    1个字节的 tinyint TINYINT
                                    2个字节 smallint
                                    3个字节 mediumint
                                    4个字节 int/integer
                                    8个字节 bigint
                                浮点数
                                    4个字节 float
                                    8个字节 double
                                    16个字节 decimal(n,m) n=>数值的位数, m=>小数的位数

                            字符串(文本/二进制)
                                文本
                                    char(n) 固定长度
                                    varchar(n) 指定范围的变长
                                    tinytext 变长 255
                                    text 变长 65535
                                    mediumtext 变长 16M
                                    longtext 变长 4G
                                二进制
                                    tinyblob 255
                                    blob 65535
                                    mediumblob 16M
                                    longblob 4G
                            时间
                                年月日
                                    date
                                时分秒
                                    time
                                年月日时分秒
                                    datetime
                                TIMESTAMP(更新操作会自动更新)

                            数组/JSON/枚举

                        修饰:
                            不允许NULL: NOT NULL， 未设置允许为NULL
                            设置默认值: DEFAULT 0 / ""
                            主键: PRIMARY KEY
                            自动增长: AUTO_INCRMENT (mysql);
                            唯一性: UNIQUE KEY

                        注释:
                            COMMENT ""
                        索引:
                            index index_name(索引列)
                            索引类型: hash/btree
                    删除表:
                        drop table tablename;
                    修改表:
                        alter table tablename add column columnName  类型 修饰 注释;
                        alter table tablename drop column columnName;
                        alter table tablename modify column columnName 类型 修饰 注释;
                    删除表数据:
                        tuncate table tableName;

                    查看当前库有哪些表
                        show table;

                    查看表结构
                        desc tableName;
                        show create table tablename;

                    索引:
                        create [unique] index indexName(column) on user;
                        drop index indexName on user;

            DML => 数据操作语言
                INSERT/UPDATE/DELETE
            DQL => 数据查询语言
                SELECT
    GO操作MySQL
        1. 找一个驱动 sql接口规范 => 驱动实现
          https://pkg.go.dev/github.com/go-sql-driver/mysql
        2. 初始化驱动 下划线导入
            github.com/go-sql-driver/mysql
        3. 导入database/sql
                对数据库操作 连接池 => 对象池 sync.Pool
        4. 连接 Open host:port user:password dbName chaset=utf8mb4 parseTime=true
        5. 测试 Ping
        6. 操作:
            DML Exec
                INSERT/DELETE/UPDATE
                最后插入的ID 受影响的函数
            DQL Query
                SELECT
                查询的结果
    练习

    用户管理
        用户:
            id 主键 自动增长
            name 字符串类型 最大长度为64 不允许为NULL
            password 字符串类型 最大长度1024 不允许为NULL
            birthday 年月日类型 不允许为NULL
            telephone 字符串类型 最大长度32
            email 字符串类型 最大长度64
            addr 字符串类型 最大长度64
            部门ID
            状态
            角色:
            createdAt 创建时间
            updatedAt 修改时间
            deletedAt 删除时间
        角色:
            id
            name
            createdAt 创建时间
            updatedAt 修改时间
            deletedAt 删除时间
        功能模块：
            id
            name
            createdAt 创建时间
            updatedAt 修改时间
            deletedAt 删除时间
        角色和模块多对多关系
        部门
            id 主键
            name 部门名称

            createdAt 创建时间
            updatedAt 修改时间
            deletedAt 删除时间

        资产:
            id
            ip
            addr

        创建表
            create table
            关系:
                1:1
                    /etc/passwd id,name
                    /etc/shadow id,password
                1:n
                    用户 和 角色
                    admin/kk        管理员/操作员

                    用户 role_id
                m:n
                    书作者  书
                    A,B, C   book1, book2
                    A B book1
                    B C book2

                    第三方表
                    book id  auth id

        添加数据:
            资产表:
            insert into asset values(1, "192.168.0.2", "北京机房", now(), now(), NULL);

            insert into asset(`ip`, `addr`, `created_at`, `updated_at`) values("192.168.0.3", "北京机房", now(), now());

            insert into asset(`ip`, `addr`, `created_at`, `updated_at`) values
            ("192.168.0.4", "北京机房", now(), now()),
            ("192.168.0.5", "北京机房", now(), now()),
            ("192.168.0.6", "北京机房", now(), now()),
            ("192.168.0.7", "北京机房", now(), now()),
            ("192.168.0.8", "北京机房", now(), now()),
            ("192.168.0.9", "北京机房", now(), now());

            删除数据
            delete from asset;
            delete from asset 条件

            更新数据
            update asset
            set addr="西安机房", updated_at = now()
            ;
            查询所有数据
                select * from asset;
                select ip, addr from asset;

                查询数据量
                select count(*) from asset;

                去重查询
                select distinct ip from asset;

                limit 限制查询数量
                offset 偏移
                分页
                limit 5 offset 0;
                limit 5 offset 5;

                排序:
                    order by name asc|desc, name asc|desc;

                过滤条件
                    where (条件)

                    条件: bool
                            布尔运算
                                与 and
                                或 or
                                非 not
                            关系运算
                                <
                                >
                                >=
                                <=
                                !=
                                =
                            字符串类型
                                like
                                    以xxx开头   like 'abc%'
                                    以xxxx结尾  like '%abc'
                                    包含xxxx    like '%abc%'

                                                like '|abc|' ESCAPE '|'
                                                _ 只有一个

                                                a?c

                                                like 'a_c';

                            数组
                                status = 1
                                status = 1 or status = 2
                                status in (1, 2)
                                status not in (1, 2)
                            between and
                                status >= 1 and status <= 10
                                status between 1 and 10
                                status not between 1 and 10

                            NULL
                                status is null;
                                status is not null;

                                deleted_at now()

                                deleted_at is null;


                        where order by limit offset
                聚合查询:
                    统计:
                    select 聚合列, 聚合函数 from tablename group by 聚合列;

                    每个ip出现次数
                    select ip, count(*) from asset group by ip;

                    select 聚合列, 聚合函数 from tablename where group by 聚合列 having ;


                    select ip, addr, count(*) from asset group by ip, addr having count(*) > 1;

                    聚合函数：
                        count
                        max
                        min
                        avg
                        sum

            多表联查
                user department
                user role

                insert into department(name, created_at, updated_at) values
                ('开发', now(), now()),
                ('运维', now(), now()),
                ('测试', now(), now());

                insert into user(name, birthday, created_at, updated_at, description, department_id) values
                ("a", now(), now(), now(), "", 1),
                ("b", now(), now(), now(), "", 1),
                ("c", now(), now(), now(), "", 2),
                ("d", now(), now(), now(), "", 0);

                select user.id, user.name, department.name
                from user, department
                where user.department_id = department.id;

                select user.id, user.name, department.name
                from user, department
                where user.department_id = department.id
                    and department.name='运维';


                left join
                right join
                inner join => join

                select user.id, user.name, department.name
                from user left join department on user.department_id = department.id;

                select user.id, user.name, department.name
                from user right join department on user.department_id = department.id;

                select user.id, user.name, department.name
                from user inner join department on user.department_id = department.id;


                select user.id, user.name, department.name
                from user join department on user.department_id = department.id;


                insert into module(name, created_at, updated_at) values
                ("用户管理", now(), now()),
                ("机房管理", now(), now()),
                ("配置管理", now(), now()),
                ("项目管理", now(), now());


                insert into role(name, created_at, updated_at) values
                ("开发", now(), now()),
                ("测试", now(), now()),
                ("管理", now(), now()),
                ("运维", now(), now());


                insert into rel_role_module(role_id, module_id, created_at, updated_at) values
                (3, 1, now(), now()),
                (3, 2, now(), now()),
                (3, 3, now(), now()),
                (3, 4, now(), now()),
                (4, 2, now(), now()),
                (4, 3, now(), now()),
                (4, 4, now(), now()),
                (2, 3, now(), now()),
                (2, 4, now(), now()),
                (1, 4, now(), now());

                select role.name, module.name
                from role,module,rel_role_module
                where role.id=rel_role_module.role_id and module.id=rel_role_module.module_id;


                select role.name, module.name
                from role inner join rel_role_module on role.id=rel_role_module.role_id
                inner join module on rel_role_module.module_id = module.id;


                select user.name, role.name, module.name,department.name
                from user,role,module,rel_role_module,department
                where user.role_id = role.id and role.id=rel_role_module.role_id and module.id=rel_role_module.module_id and user.department_id=department.id;

            子查询:
            select where select

            想要查询部门名称中包含 '测试' 所有的人
            select * from user where user.department_id in (
                select id from department where name like '%测试%'
            )

            别名
            select r.name ruleName, m.name as moduleName
            from role r,module m,rel_role_module as rel
            where r.id=rel.role_id and m.id=rel.module_id;

            集合
                并集 union/ union all

                select name from role
                union all
                select name from department;

                交集 intersect

                select name from `role` intersect select name from `department`;

                差集 except

                    select name from `role` except select name from `department`;

            函数
                字符串:
                    upper()
                    lower()
                    trim()
                    to_base64
                    from_base64
                时间:
                    now
                    date_format
                编解码:
                    md5
                    sha1

insert into user(name, password, birthday, created_at, updated_at, description, department_id) values
                ("kk2", "kk@123", now(), now(), now(), "", 1),
                ("admin", "kk@123", now(), now(), now(), "", 1);
