1. promagentd实现
    a. 注册
    b. 心跳
    c. 获取配置
        server api
            uuid version

            db version > request version
2. cmdb server
   a. 注册/心跳 实现
        web 展示终端信息 和心跳时间(正常/不正常)

        |主机名(正常(绿色)/不正常(红色))| 服务器IP(Addr) | 配置版本 | 更新时间 | 操作|
    b. 可以删除 (逻辑删除)
        deletedat = now()

    c. 编辑
        只编辑 配置(version = time.Now())