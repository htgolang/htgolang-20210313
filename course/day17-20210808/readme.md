1. 复习(10分钟)
    prometheus

    采集数据
        不可修改源码:
            exporter
        自主研发:
            嵌入应用程序

    client_golang
    yaml
3. 作业
4. 新内容
    exporter mysql
    cmdb prometheus metrics
        Filter Before
            总请求次数
                Counter
            每个URL请求次数
                Counter

        Filter After
            每个状态码的出现次数
                Counter
            每次URL请求延迟时间统计
                Histogram



    prometheus配置管理
        cmdb web prometheus配置
        promagentd
            API:
                注册
                心跳
                配置获取
            Agentd:
                注册
                    周期性 1h
                心跳
                    周期性 1m
                获取配置
                    周期性 1m (配置未变化 不更新)
                    version/updated_at

                    prometheus.yml
                    kill -HUP pid

    promagent 终端表的结构
    ID
    uuid
    hostname
    addr

    config
    config version

    heartbeat_time

    created_at
    updated_at
    deleted_at