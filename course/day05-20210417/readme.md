1. 复习
   5-10分钟简单会议上节课内容
   + 同学们主动复习 => 一定要复习
        重要节点: 第8节课, 第12节课, 第15节课
   + 留出更多时间上新内容

    day04:
        包/模块
            代码组织 => 项目组织方式
            包：
                定义
                    package pkgname;
                    每个文件夹内包名必须一致
                使用
                    import 包路径;
                    绝对导入
                    下划线导入
                    别名导入
                    相对导入 不要使用 go module禁用
                    点导入
            模块:
                go mod init github.com/imsilence/utils
                go mod tidy
                go run
                go build
        标准包 => 自己学习
        第三方包
            getpass => 从控制台接收密码
            github.com/howeyc/gopass
        单元测试
            单元测试用例
            测试覆盖率
            基准测试

2. 作业
    已交作业12人: 坚持
    未交作业的同学: 作业补上(给助教说下, 以方便助教进行评阅)

3. 新内容
    自定义类型:
        + 想要自己定义一个由不同类型元素(属性，字段)组成一个复合类型
        + 将现有的类型进行重新定义, 对新的类型提供一些方法, 与原有类型类型无关, 具有原有类型的特性, 但是可以与原有类型进行相互转换. 计数功能  int 计数类型
            别名&重新定义
            - 别名, 与原有类型相同，无区别， 不能定义新的方法
    自定义类型:
        var
        const
        func
        type: type TypeName Formatter
        Formatter:
            重定义: type Counter int
            结构体: type User struct {
                属性名称1 属性类型1
                属性名称2 属性类型2
                ...
                属性名称n 属性类型n
            }
    别名:
        type TypeName = Formatter

   + 面向对象: 思想=>实现, 三大特征: 封装，继承，多态
     + 结构体
     + 方法


    new,make
        make => 创建对象
            slice,map,chan
        new => 返回指针 创建结构体指针对象
            new 基本数据类型
            var i int = 1
            p := &i

            p := new(int)

            new(Struct) <=> &Struct{} => New


            // 值类型 => var u User => 可以使用 非nil
            // 指针类型 => var u *User => 不可以使用 nil

            var a *int = new()
            var a *int
            a = new()

            var u *User = new(User)
            u = new(User)
            u := new(User)

    1. 定义变量 => 多个相同数据类型 => 合并
    type user struct {
        id int
        name string
        addr string
        tel string
    }

    type user struct {
        id int
        name, addr, tel string
    }

    调用结构体，调用函数 参数很多, 一行写一个参数(属性值), 注意函数，结构体结尾括号位置 与最后一个元素之间的逗号， 再不同行不能省略

    结构体的可见性

    面向对象 代码复用=> 继承、组合
        嵌入 => 组合模式

    已经有了多个结构体, 在当前结构体中使用已经实现的结构体
    购物系统
        商品商家 地址
    User
        地址

    Addr
        省份
        城市
        街道

    方法: 给特定类型定义的函数，只能由特定的类型进行调用
        func funcName(args)returns {

        }

        接收者 => 指定函数属于哪个类型，只能由对应的类型进行调用
        func (reciver) methodName(args) returns {

        }

        除了接收者，方法的参数/返回值 与go中的函数完全一致

        接收者也是一个参数: t T 参数类型 方法所属的类型, 方法只能由T类型的对象进行调用
        T 自定义类型
        t 变量名, 在对象调用方法时传递调用的对象

        1. 值接收者: t T
        2. 指针接收者: t *T
        3. 方法之内不需要访问/修改t对象的属性, t省略 T

        可见性: 与属性相同
            首字母大写 => 包外可见
            首字母小写 => 包内可见

        嵌入 => 组合