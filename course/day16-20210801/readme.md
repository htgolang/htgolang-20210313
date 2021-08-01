1. 复习 10分钟
    beego
        组件:
            cache
            session
            log
            tasks
        MVC:
            C: Namespace
               Filter
            ORM: 关系
        配置:
            https
    golang：
        regex
        cobra
        errors
        Stringer接口
            String

    部署：
        发布：制作包(编译->打包) -> 部署(拷贝->解压->修改配置->启动)
        进程监控: supvisor, systemd
2. 作业
3. 新内容
   a. Prometheus
        0: 没接触/听过
        1: 使用过(搭建过)
        2: 深度使用(自己部门在使用，维护，基于prometheus与开发功能)

        基本原理
        部署
            node exporter: 数据提供
            prometheus server: 数据处理 和 存储
            alertmanager: 告警的处理者

            ---
            yaml(yml): 类似与json 格式化的配置文件
            {key:value, key, value}
            key: value
            value 类型:
                字符串
                    key: "x xx"
                    key: xxx
                数值
                    key: 1
                    key: 1.1
                boolean
                    key: true
                    key: false
                数组
                    key: [1, 2, 3]

                    key:
                    - 1
                    - 2
                    - 3
                映射
                    key: {attr1: value1, attr2:value2}
                    key:
                        attr1: value
                        attr2: value

            认证：
                web basic
                    user:password

                prometheus scrapy:
                    Web Basic
                    Token方式

                    Authorization: Token
                        web basic: Basic[空格]base64(user:password)
                        bearer token: Bearer[空格]Token

                            格式 token jwttoken

                    target 配置认证方式
                        应用 exporter

                    scrapy 配置信息
                        prometheus
                            web basic: user:password

            1s 100
            2s 100

            满足50%  1s

        操作
        接口
            node_cpu_total
            count(node_cpu_seconds_total{mode="system"}) by (instance)
            sum(100 * (1-irate(node_cpu_seconds_total{mode="idle"}[5m]))) by (instance)


            node_memory_MemTotal_bytes

            100 - (node_memory_MemFree_bytes + node_memory_Buffers_bytes + node_memory_Cached_bytes + node_memory_SReclaimable_bytes) * 100 / node_memory_MemTotal_bytes

            (1-(node_memory_MemFree_bytes + node_memory_Buffers_bytes + node_memory_Cached_bytes + node_memory_SReclaimable_bytes) / node_memory_MemTotal_bytes) * 100



    b. 提供数据
        指标类型
            Counter
            Guage
            Historgram
            Summary
        根据标签是否未动态
            固定标签
            动态标签

        开发流程
            1. 指定指标
            2. 注册指标
            3. 启动web服务器(注册promhttp Handler)

        指标信息更新:
            指标信息提供者
               1. 事件触发型
                   http请求时更新
               2. 时间触发型
                   周期性更新
            prometheus 调用metrics api 触发获取

        web.basic:
            Authorization: Basic base64(xxx)


        go 库：
            信号:
            日志:
            yml: