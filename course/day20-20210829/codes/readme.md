今天内容
1. osquery插件
    注册
        register

        {
            uuid
            hostname
            addr

        }
    心跳
        heartbeat
            uuid
    配置获取更新
        config
            uuid configversion
    任务获取 执行 结果上传
        task
            uuid
        result
            uuid
            status
            list
    日志上传
        logs
            uuid
            list

2. 任务下发(结果)

3. gin通信


日志 放在二进制程序相同目录
    /opt/agentd/agentd

    ./logs/agentd.log 相对路径 程序启动的目录
        cd /tmp => /tmp/logs/agentd.log
        cd /home/kk => /home/kk/logs/agentd.log

    获取程序的路径 所在目录

    os.Args