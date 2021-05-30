1. 复习 10分钟
   并发编程
        主例程
        工作例程 需要保证正常结束 waitgroup
                 不需要保证正常结束
        go func() {}()
    并发通信
        共享数据+锁
            sync.Mutex
        管道 CSP
            chan T
            nil => 一定要初始化后才能使用
            make(chan T)
            make(chan T, bufferSize)

            读/写/关闭
            value <- channel
            value, ok <- channel
            channel <- value
            遍历
            for v := range channel {

            }

        select case

        context 正常结束一个例程
        time包
        sync包
        runtime
        socket开发 服务端/客户端
2. 作业处理
    GO语法已经全部学些完成 => 可以看使用go语言写的任何功能代码
    Go语法+业务逻辑(站在当时写的人的思维角度)

    泛型
    集合 slice T, map T
    针对操作: 函数 slice T, map T
3. 新内容
    socket开发
        图解TCP/IP

        网络 + GO开发网络功能(net)
        分层模型 OSI
        TCP/UDP协议
        我要开发什么东西: 服务端 客户端
        如何开发

        TCP:
            服务端
                1. 创建监听
                2. for
                   a. 接收客户端连接
                   b. go 处理客户端业务
                        [
                            规则一：读/写顺序、次数
                        ]
                        获取数据(读数据)
                        响应数据(写数据)
                        [
                            规则二: 数据格式
                        ]
                3. 关闭监听
            客户端
                1. 连接服务
                2. 处理业务
                   a. 获取数据
                   b. 响应数据
                3. 关闭连接

    web开发
        图解HTTP

        c/s => client + server
        web => b/s => browser + server

        WEB API = 第三方使用者，使用目标不是浏览器
                  需要自己去调用

                  组装HTTP请求 发送给服务端 解析HTTP响应



        客户端  <=HTTP=>服务端

        HTTP协议: 超文本传输协议
            HTTP 1.1/1.0/2.0 => 3.0
            HTTP 1.1/1.0 文本

            请求和响应的文本格式

            请求都是从客户端发起(请求/应答)
            HTTP Request
                \r\n分的多行文本数据
                1: 请求行 第一行 Method URL Protocol/Version
                    Method: 请求方式
                            OPTIONS
                            HEAD GET POST DELETE PUT
                            Connect TRACE

                            针对浏览器常发起的请求有GET, POST
                    URL: 标识不同的服务
                        针对web开发需要定义URL和URL处理逻辑
                2: 请求头 多行 Key: Value
                        Host: 服务主机名
                        浏览器信息 UserAgent
                        会话信息 Cookie
                        Content-Type:
                        ...
                [空行]
                3. 请求体 可能没有 如果有内容 格式
                    常用编码方式
                        application/x-www-form-urlencoded
                            key=value&key=value
                        multipart/form-data
                            上传文件
                        application/json
                        application/xml

            HTTP Response
                1. 响应行
                   协议 响应状态码  响应状态码文本描述
                   1xx 告知客户端数据已收到请继续上传数据
                   2xx 请求正常
                        200
                   3xx 请求重定向
                        301 永远
                        302 临时重定向
                   4xx 请求数据错误
                        400 提交数据错误
                        401 认证错误
                        402
                        403 权限拒绝
                        404 URL错误
                        405 请求方式错误
                   5xx 服务端错误
                        500
                        501
                        503

                2. 响应头
                    key: value
                    Content-Type: 响应格式
                    Set-Cookie: 设置会话
                [空行]
                3. 响应体
                    text/html
                    application/json
                    application/xml

        如何开发：服务端 客户端 net.http包
            url => http.Handler 处理器
                    ServeHTTP(http.ResponseWrite, *http.Request)

            客户向服务器提交数据:
                url?key=value&key=value
                queryparams =>

                request body:
                    POST
                    content-type: application/x-www-form-urlencoded
                                    key=value&key=value
                                  application/json
                                  multipart/form-data

        gin框架

    rpc开发
        服务端 客户端
        grpc: 在原有内容基础之上会给大家额外添加，需要占用1个小时左右