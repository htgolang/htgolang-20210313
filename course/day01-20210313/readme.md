0. 调查
    有编程经验: 1, python,go,java,shell
    无任何经验: 0

    编程语言是用来干嘛的？
        a. 告诉计算机如何完成工作 => 解决问题思路
        b. 给人看的 => 代码的规范(风格)

    编程语言是解决问题的工具, 选择自己合适的先去学习流程，并解决问题

1. 课程目标
    养成编程的思路
    学习GO语言
    完成项目 在工作中使用
2. 课程要求
    集中精力，及时答复
    课堂代码 => 自己敲一遍
    上课节奏以老师为准
    互动
    课后 练习 => 一定要完成
3. 学习方法

4. 前提准备
    GITHUB: 上课练习、资料、代码
        https://github.com/htgolang/htgolang-20210313
        大家需要准备的:
        a. 注册github账号
        b. 加入群组
            自助加入: http://118.24.20.160:8091/
        c. git使用
            + 链接：https://pan.baidu.com/s/1yBSDO_4GV0gpKf7QiYFJCw 提取码：o1t5
            + 链接：https://pan.baidu.com/s/1Okj6yoRi9ud1tUG_niK82g 提取码：yus3

    工具：
        操作系统：window10, centos7.8
        虚拟机: hyper-v/vmware
        远程控制: xshell, lrzsz(sftp)
        代码： vscode => 禁止用vim
        go: 1.16
        mindmanager: => html

5. 今日课程
    初识
        go 编译型语言 静态的 命令式的
        go 源码 => 编译器 => 二进制程序

        a. 编译器 => 源码 -> 二进制
        b. 编辑器 => 写源码
    
        go环境安装
            a. 安装
                windows: msi
                linux: tar.gz
                        解压 /usr/local
                        tar zvxf .gz -C /usr/local
                        配置环境变量
                        go $PATH
                        GOROOT=>go安装路径 /usr/local/go
                        GOPATH=>$HOME/go
            b. 卸载
                windows: 控制面板卸载,配置信息删掉,下载文件删掉
                linux: 删除文件/删除环境变量配置
    
            配置:
                GO111MODULE=on
                GOPROXY=https://goproxy.cn,direct
    
                a. 配置环境变量
                b. go env -w
    
        第一个程序:
            notepad:
    
        vscode:
            安装
            想要有提示功能 go插件/ctrl+shift+p -> go install/update tools
            termianl: windows => cmd(对powershell不熟悉的用cmd)

    问问题方式：
        1. 概念不清楚
            描述清楚问题
        2. 报错
            描述清楚问题
            代码截图
            报错截图

    计算圆形的面积 半径 5
    Πr平方 3.1415926 * 5 * 5

    GO基础
        基本组成单位
            标识符
                命名规范
                a. 组成元素
                    非空的Unicode字符 数字 下划线
                b. 不能以数字开头
                c. 不能使用go中的关键字

                d. 不能定义空白标识符 _, 仅使用
    
                e. 避免于go中的内置标识符冲突
                f. 仅使用英文字母，数字，下划线定义标识符
                    名字, 半径
                g. 变量命名 多个英文单词 驼峰式
                    add user
                    addUser 驼峰式 go推荐
                    add_user 下划线式
                h. 简明知义

    声明
        定义标识符

        变量声明 var
            var 标识符 类型
            变量定义的位置:
                go文件中最外围
                定义在函数内
        常量声明 const
        函数声明 func
        自定类型声明 type

    作用域
        显示声明作用域 {}
        // 隐式的作用域
            全局作用域
            包作用域
            文件作用域
            if，switch, for, select case

    基础数据类型
        布尔类型
            含义: 表示真假
            类型名(标识符): bool
            零值: false
            内存占用： 字节
            值: true(真)/false(假)
            操作:
                逻辑运算
                    与 &&: expr1 && expr2
                        只有两个操作数为true结果true
                    或 ||: expr1 || expr2
                        只要有一个为true结果为true
                    非 !: !expr
                        expr为true，结果false
                        expr为false 结果true
                关系运算
                    等于 == 不等于 !=
            格式化输出:
                fmt.Printf 占位符 %t

        数字类型
            整数
                1字节 0000 0000
                有符号: 可正/可负, 首位: 1负, 0正
                        000 0000 -2 ^ (n-1) ~ 2 ^ (n-1) - 1
                                -128 - 127
                无符号: 只表示>0
                        0000 0000
                        0 ~ 2 ^ n - 1
                        0 ~ 255
    
                int/uint //32位=>4字节，64=>8字节
                byte 1字节 字节类型 ascii字符 -> 整数
                rune 4字节 码点 unicode编码 字符->整数
                int8/uint8 1个字节
                int16/uint16 2字节
                int32/unit32 4字节
                int64/uint64 8字节
                uintptr //32位=>4字节，64=>8字节
    
                零值: 0
                字面量:
                    x * base ^ (n -1) 求和
                    十进制: 10, 1, -1, -100
                    八进制: 0xx(x< 8) 010 - 1 * 8 + 0 = 8
                    十六进制: 0Xmmm(m<16) 0-9 10-15 a-f 0xf => 15 0x1f=>1 * 16 + f
                    二进制: 0bxxxx(x=0,1) 0b1111 => 8
            操作:
                算术运算：
                    +
                    -
                    *
                    /
                    %
                    ++
                    --
                关系运算: bool
                    >
                    >=
                    <
                    <=
                    !=
                    ==
                位运算: int 10-> int 2
                    &: 两个都位1 结果位1
                    |: 只要1个位1，结果位1
                    ^(单目): 1->0, 0->1
                    ^(双目): 相同 0 不同 1
                    <<: 左边移动n，补0
                    >>: 右边移动n位，补符号位
                    &^: 位清除, 右操作数位1->0, 0->保持不变
                赋值:
                    =
                    +=
                    -=
                    *=
                    /=
                    %=
                    &=
                    |=
                    ^=
                    <<=
                    >>=
    
                    a += b => a = a + b
                    a &= b => a = a & b
            浮点数(小数)
                类型名：
                    float32 -> 32位(4字节)
                    float64 -> 64位(8字节)
                零值: 0.0 => 0
                字面量:
                    十进制的小数: 10.01
                    科学计数法: a * 10 ^n
                            a e n
                            1000 => 1 * 10 ^ 3
                            1e3
    
                操作：
                    算数运算
                        +
                        -
                        *
                        /
                        ++
                        --
                    关系运算
                        >
                        <
                        <=
                        >=
                        等于/不等于 |a-b|<精度
                    赋值运算
                        =
                        +=
                        -=
                        *=
                        /=
                强制类型转换
                    float32()
                    float64()
                格式化输出
                    %f
                    %e
            复数
                i*i = -1
        字符串类型
            string
            零值: ""
            字面量:
                可解析字符串： "xxxx"
                原生字符串: `xxxx`
    
            特殊字符
                \t tab
                \r 回车
                \n 换行
                \a 响铃
                \b 退格
                \f 换页
            转义: \
    
            操作：
                字符串连接: +
                关系运算:
                    >
                    >=
                    <
                    <=
                    ==
                    !=
                    从左边第一个元素开始`比较(byte) 能确定大小则结束，则比较下一个字节
                    "abz" > "ac" => false`
                赋值运算
                    =
                    +=
                索引操作 => byte
                切片操作 => byte
    
        指针类型
    流程控制