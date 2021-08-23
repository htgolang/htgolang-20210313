1. 复习
    prometheus
        agentd
            -> server 拉去配置 -> 检查 -> 更新

        yaml

        server
            -> web 富文本编辑器(yaml)->检查->更新数据库

        web配置->agentd没有拉取到(没有更新成功)
2. 作业
3. 新内容
    主机安全的项目

        配置
            内网IP地址范围
            常用端口
            扫描周期

        扫描：
            手动扫描
            后台定时扫描

            同时运行一个
                lock + 标识

        展示：
            ip, 描述, 状态, 时间

        有人偷偷上线一个系统：有
        有人上线 服务暴露到外网: redis/es

        IPAddr: 192.168.0.0-192.168.255.255

            10.0.1.0;192.168.0.0-192.168.255.255; , 10.0.0.0/24

            ;,\s

            ip格式化

            [][start, end] ports
                192.168.0.0, 192.168.255.255
                10.0.0.0, 10.0.0.255

        主机扫描
        端口扫描

            masscan: 主机扫描
            nmap: 主机扫描，端口端口，服务扫描

            exec执行系统命令 => 文件(xml/json/text) => 解析 => 结果

        TCP 全连接扫描
            TCP client = connect => TCP server
                ok => 在线
                error => 不在线

                ip:
                port: 常用端口列表: 22, 3389, 80, 443, 8080, 8888, 8443, 21, 6379, 3306

        资产管理：
            ip唯一标识
                service 对外暴露的端口
                    端口扫描 0-1024, 0-65535

        prometheus监控
            web metrics

            smtp
            redis
                connect server
                set testkk 1;
            mysql
                exporter -> connect server Ping
                            ok => 正常
                            error => 异常

                select 1;

        监控 网站/api 故障，延迟
            http client 发起请求
            http response status/content


        主机威胁检查
            osquery

        gin


        0. 调用osquery查询信息
        1. osquery table
            自定义表
        2. config插件
        3. log插件
        4. 分布式插件