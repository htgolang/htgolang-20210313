1. 复习(10分钟)
   prometheus agent
        注册 => Server Agent手工添加
        心跳 => Server 心跳时间
        配置获取 =>
            Server 编辑配置
            Agent 感知配置变化 version
            检查配置
            应用配置
    Server:
        Agent列表展示
        配置编辑
        Agent删除
            agent 123 deleted_at = now

    第三方库
        logrus
        lumberjack
        viper
        req

    标准库
        signal

2. 作业
3. 新内容
    a. prometheus配置集中管理
        cmdb server
            mvc
                controller (form)
                service
                model
                view
        prometheus agentd
            go
    b. 告警管理
        alartmanager -> http api

        cmdb server 告警接收
        告警查询/处理
            分页

    form:
        toModel => form attribute => model
        FillModel => model => form attribute

    1. 编辑
        样式展示
            codemirror
        编辑格式校验
            yaml Unmarshal
        内容校验

    authorzation:
        type: Token
        credentials: b

    authorization: type credenctials(a b)


    {
        "status": "firing",
        "labels": {
            "alertname": "Node CPU Percent is High",
            "env": "dev",
            "instance": "localhost:9102",
            "severity": "high"
        },
        "annotations": {
            "description": "Node localhost:9102 CPU Percent is High 200",
            "summary": "Node localhost:9102 CPU Percent is High"
        },
        "startsAt": "2021-08-15T09:20:26.851Z",
        "endsAt": "0001-01-01T00:00:00Z",
        "generatorURL": "http://centos:9090/graph?g0.expr=node_cpu_percent+%3E%3D+80\u0026g0.tab=1",
        "fingerprint": "bd97c3c7d278748a"
    }

    id:
    reason:

    fingerprint: varchar
    name:   varchar
    status: 0/1

    description: varchar

    startsAt: *time.Time
    endsAt: *time.Time resolved

    generatorURL varchar

    labels: text
    annotations: text

    1. firing
        数据库中添加记录就行了
        a cpu 高
        a cpu 高

        fingerprint 查询 有没有相同的告警 未解决的
            有 不用存储
            无 增加
    2. resolved:
        查询 fingerprint 查询 有相同的告警 未解决的
            有 更新状态解决  endsAt