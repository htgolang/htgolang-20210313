1. 复习(10分钟)
    Beego web开发框架 全栈的
        MVC设计模式的实现
            Controller
            View
            Model
        ORM设计模式的实现
    Beego 项目的结构:
        conf 配置
            app.conf
        controllers 控制器
        routers 路由配置 url => controller
        services 服务
        models 模型
        views 视图
        tests
        static 静态资源文件
        forms request数据
        main.go 程序启动入口

    Beego服务启动
    Beego配置
    路由+控制器
        Controller => web.Controller
        接收用户数据
        响应
            使用模板:
                TplName = ""
                Data[xxx] = xxx
            JSON:
                Data[json] = xxx
                ServeJSON()
        Session控制
        验证
        url构建 url -> Controller => controllerName.actionName -> url
            template urlfor
            go代码 web.URLFor
        flash消息
            利用cookie
        错误处理

2. 作业
3. 新内容
    Beego ORM
    ORM: 对象关系映射 Object Relation Mapping
        关系数据库          <=> 面向对象类
        表(定义)            <=> 结构体(类)
            名字            <=> 结构体名(类名)
            列              <=> 属性
                列名        <=>     属性名
                列类型      <=>     属性类型

            示例:
                go 对象 <=> 反射 => create table
                表结构 => 结构体

        数据:
            每行数据        <=> 实例化/类对象
        数据操作:
            增删改查        <=> 方法(自动操作转换为SQL 调用数据库执行SQL语句)

    ORM思想 => 实现beego orm, gorm, ...(工具)
        orm => 针对不同数据库进行实现
        优势:
            1. 如果在使用过程中未使用某个数据库特性的SQL或者或针对某个数据库的特性功能，可以实现数据库之间的切换
            2. 可以在不了解SQL的情况下 实现对数据库操作
        缺点:
            1. 性能
            2. 只能使用ORM框架提供的基本功能，针对数据库提供的特性功能，ORM未实现的功能 只能使用原生SQL

    使用:
        使用包:
            数据库驱动 github.com/go-sql-driver/mysql
            orm库: github.com/beego/beego/v2/client/orm 使用database/sql并提供orm的增强功能

        定义结构
            表结构 = sql
            可以不定义表结构 => 可通过orm同步
                创建表          =>          同步
                删除表          =>          不同步
                对于表列修改
                    添加列      =>          同步
                    修改列      =>          同步
                        列名    =>          可同步, 新增
                        列的属性 =>         可同步(beego不会同步)
                    删除列      =>          同步

            需要定义结构体(类)
                type User struct {
                    Id int64
                    Name string
                }
                需要定义结构体和表的关系
                    表关系
                        表名 <=> 默认 结构体名驼峰式 -> 全小写下划线式
                                除首字母外, 碰到大写字母加下划线转小写
                                 User <=> user
                                 UserDB <=> user_d_b 可能user_db
                                 需要可定义与表关系 => 方法
                        修饰
                            索引 => 方法

                    列的关系
                        列名 <=>  默认 结构体属性名驼峰式 -> 全小写下划线式
                                 属性标签
                        修饰 <=> 属性标签

                        标签名: orm

            告知ORM管理的结构体(注册)
                orm.RegisterModel

            表结构:
                同步

            表数据
                增
                删
                改
                查

        问题：
            可以现通过sql创建表后再通过使用orm操作吗
            如果用原生SQL运行一段时间的业务(已经有数据了)可以修改为使用orm操作

        beego orm模块可单独使用或与其他第三方模块结合使用
            单独使用orm对数据库增删改查
            和http，gin包结合对数据库操作

        模型定义:
            表
                表名
                索引
            属性
                属性名
                属性修饰

                标签: orm
                标签项
                    pk
                    auto
                    index/unique
                    column
                    type
                        int int
                        int64 bigint
                        string varchar(255)
                            size()
                            type() text(longtext)/char
                        time.Time *time.Time datetime
                            type(date)
                    null
                    default
                    description

                    针对小数类型
                        digits
                        decimals

                    针对时间 *time.Time:
                        auto_now
                        auto_now_add
