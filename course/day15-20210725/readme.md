1. 复习(10分钟)
    MVC => view
    beego => view => html/template
        函数定义和使用
        复用:
            Layout
            LayoutSections
    html
    css:
        基本知识
        位置:
            内联
            页面内
            页面外
        绑定(选择器)
            标签 tagName
            ID #id
            class .className
            属性选择器 [attrName]
                        [attrName=attrValue]
            子选择器
                    parentSelector > childrenSelector
            子孙选择器
                    parentSelector childrenSelector
            复合选择器
                tagName.className
        css定义:
            key: value

            定位: position
            盒子模型:
                margin+border+padding+(width/height)

    bootstrap:
        1. 设计好的样式 直接使用
        2. link static 使用 指定标签的class form table
        3. 布局 + 栅格模型
2. 作业处理
3. 新内容
    a. beego
        1). 配置
            app.conf
                runmode="dev"
                httpport=8888

                include "dev.conf"
                include "prod.conf"

            dev.conf
                [dev]
                httpport=9999

            prod.conf
                [prod]
                httpport=8080

            mysql.conf
                [mysql]
                host=xxxx

            redis.conf
                [redis]
                host=xxx

        kill -n pid
        进程 => 代码 对信号进行处理的


        2). 路由 Namespace
            controller/action
            controller  method get/post


            controllers
                users
                    IndexController
                        Index
                roles
                    IndexController
                        Index

            web.AutoRouter(&users.IndexController{})
            web.AutoRouter(&roles.IndexController{})

            /index/index => ?Controller
            /users/index/index => users.IndexController
            /roles/index/index => roles.IndexController

            Namespace

        3). 日志

        4). csrf
            API => 第三方提供服务 不开启csrf检查

            前后端分离 => token => js

            Controller
                template => csrf

                http response
                http response header csrf
                    set-cookie

                ajax => post
                        http request header(自定义的header不能用cookie)

            框架
                controller:
                    token
                        session 后端服务存储

                        cookie token+sign
                view:
                    template token => form表单
                    js 解析cookie => token

                    提交时候 cookie token:sign

                服务端验证:
                    cookie token: sign => 检查token
                        header/form => xsrftoken
        5). 缓冲
            cache[key] => value

            先找缓冲 => 从缓冲中拿去数据
                    无 => 计算/查询 放入缓冲 ttl

    b. go 标准/第三方库
        正则
        error

        cobra
        logrus
    c. 部署
        http/https部署
        go 交叉编译
            纯GO语言开发的库
            CGO
        web + nginx
            nginx + tomcat
            nginx + springboot
            nginx + uwsgi + django
            nginx + fastcgi + php

            go
            nginx(代理) + go

            nohup ./cmdb >/dev/null 2>&1 &
            进程管理工具
                systemd
                supervisor

                1. 执行的命令
                2. 执行命令对应用户
                3. 工作目录
                4. 退出 是否自动启动
                5. 启动时间/延迟时间
                6. 环境变量

            nginx + n * cmdb.exe

            nginx 代理 负载均衡
                session file


                    A => file
        nginx
                    B =x => 失败


            upstream
                ip:port
                ip:port

                session 不要与应用程序耦合 => 第三方存储
                    db
                    redis
                    memcache

                    CREATE TABLE `session` (
                        `session_key` char(64) NOT NULL,
                        `session_data` blob,
                        `session_expiry` int(11) unsigned NOT NULL,
                        PRIMARY KEY (`session_key`)
                    ) ENGINE=Innodb DEFAULT CHARSET=utf8;


        session: 持久化 => 编码 Gob
                gob.Register()
                任何东西都能放
                放尽量少的内容

        每个用户只能同时在一个位置登陆
            业务逻辑
                登陆 => session => id

    应用程序配置
        app.conf => 运维

        配置和应用程序分离
            自动化
                程序部署自动化部署
                app.conf自动生成
            环境变量
                监控程序 .env key=value

    http/https部署:
        自签名证书

    d. 交付
        开发 -> 测试 -> 运维
        源码交付
        安装包
        虚拟机镜像
        docker镜像 v
    e. redis