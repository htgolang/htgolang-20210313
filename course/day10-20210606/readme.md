1. 复习
    socket
        客户端
        服务端
    web
        url => 处理器/处理器函数
        启动web服务

        ResponseWriter, Request

        接收用户提交数据 Request
            url
            body
                applicaition/x-www-form-urlencoded
                multipart/form-data
                application/json

        处理
        给用户响应 ResponseWriter
    rpc
        grpc
2. 作业
3. 新内容
    web
    html
        响应
        html => 超文本标记语言
             => 由浏览器解析 html js css
        类似于xml
            文本头: <? xml version="1.0" encoding="utf-8" ?>
            标签：
                双标签： 有开始有结尾 <tagName> <childTagName> </childTagName> </tagName>
                单标签: <tagName />
            标签的属性:
                <tagName key="value" key="value" ></tagName>
                <tagName key="value" key="value"/>
        html学习什么呢?
            html规范 => 定义tagName => 应该展示为什么样子 => 浏览器厂商
                        input => 显示为输入框

            需要知道常用标签 以及展示的效果
        html 4.x
        html 5
        html:
            文件命名: .html
            头: <!DOCTYPE html>
            根标签：<html>
                        <head>
                            <!--
                                注释
                                告知浏览器的信息 比如文件编码方式 title
                            -->
                            <meta charset="utf-8" />
                            <title>我的页面</title>
                        </head>
                        <body>
                            <!-- 使用者 告诉浏览器 给用户进行展示-->
                            xxxx
                        </body>
                    </html>

            使用浏览器向服务器端发起请求的方式都有哪些，请求Method
                1. 浏览器地址栏输入地址，按回车 => GET => url?queryparams
                2. 通过html标签定义
                   a. 超链接 GET url?queryparams
                   b. 表单 method=>POST/GET
                        GET => ?queryparams 在action中是否可以指定queryparams可以
                        POST => body 在action中是否可以指定queryparams 可以
                   c. img/link/script => GET
                3. 与服务器结合
                    重定向: GET  => url?queryparams

    模板基础 => 用户管理的项目
    session/cookie => 认证 => 简单的session => 服务端保存认证信息
        cookie          session
        客户端              服务端
        第一次方式的时候会生成sid
            如何判断第一次 : Cookie 中sid是否有效，无效则认为第一次
                        sid 为空
                        sid 对应文件不存在
                        sid 对应文件已超过时间限制


            Set-Cookie

        sid                 sid=>一块存储=>file
            Cookie sid=xxxx
                             jwt  => 客户端保存认证信息

        客户端保存状态信息:
            session => file
            session => cookie => 篡改 数据签名
            cookie => value => hs256 token
    爬虫 => goquery => web 1.0开发模式 => html