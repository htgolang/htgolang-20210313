1. 复习(10分钟)
    ORM:
        调试
            orm.Debug = true
        a. 引入包
            驱动包: github.com/go-sql-driver/mysql
            ORM: githu.com/beego/beego/v2/client/orm
        b. 定义模型(指定与数据库表的关系)
            模型(名称 属性) ->表(名称 属性)
            模型属性(名称 类型) -> 表字段(名称 类型)

            Go模型结构体
                表:
                    名称 => 结构体名称 驼峰转下划线小写
                            TableName
                    索引 => TableIndex/TableUnique
                字段:
                    名称 => 属性名 驼峰转下划线小写
                            标签: orm:"column(column);";
                    修饰
                        类型
                            语言中类型 => 数据库字段类型
                            string:
                                type(varchar); size(size) => varchar(size)
                                type(text); => longtext
                            time.Time
                                type(date); => date
                                type(datetime) => datetime
                            decimal:
                                digits();
                                decimals();
                        主键
                            pk;
                        自动增长:
                            auto;
                        唯一
                            unique
                        索引
                            index;
                        是否允许为空
                            默认不允许为空
                            null;
                        默认值
                            default(value);
                        注释
                            decription(comment);

                        时间:
                            auto_now;
                            auto_now_add;
            c. 注册
                模型 RegisterModel
                数据库驱动 RegisterDriver
                数据库连接池 RegsiterDataBase

            d. 表结构同步(ddl)
                RunSyncDb(数据库, 是否删除表, 详情)
                只同步增加的字段和表(修改结构体名称,修改属性名称)

            e. 表数据操作(dql, dml)
                ormer = orm.NewOrm()
                CURD
                    ormer.Insert()
                    ormer.Update()
                    ormer.Read()
                    ormer.Delete()
                数据集(批次)
                    queryset := ormer.QueryTable(model)
                    条件(where):
                        Filter()
                        Exclude()
                        SetCond()
                            NewCondtion()

                        key__condition

                        condition:
                            等于 => (i)exact
                            开头 => (i)startswith
                            结尾 => (i)endswith
                            包含 => (i)contains
                            关系 => lt lte gt gte
                            is null => isnull
                            in => in
                    All()
                        排序
                            OrderBy("-column")
                        分页
                            Limit
                            Offset
                    Count()
                    Update
                    Delete()

2. 作业
3. 新内容
    MVC
    V beego view + css(Bootstrap)
    Beego:
            日志
        缓冲


    css:
        层叠样式表单
        css 决定标签在浏览器的样子

        模特 =>
            人 => HTML
            衣服 => CSS
            姿势 => JavaScript

        鸡 =>
            鸡(拔完毛) => HTML
            羽毛 => CSS
            展翅 => Javascript

        css:
            设置方式:
                内联
                    直接指定在标签中
                    div 宽度 高度 width heigh
                    style=""
                页面内选择
                    head/style => 选择器(css样式 关联到 html标签)
                页面外选择
                    css单独定义到css文件中
                    html引入 css
                        head/link
                            href=>css文件路径
                            rel=> stylesheet
                            type="text/css"
            选择器:
                查找html标签
                基本的选择器
                    ID
                        查找某个标签, #id

                        <tag id="id"></tag>

                    TagName
                        查找某相同标签名的标签, tagName
                        a
                        div

                    Class
                        通过标签中class值进行选择, .className
                        <tag class="className1 className2">

                    复合:
                        tagName.className

                其他选择器
                    属性选择器
                        [attrName]
                            [href]
                                a
                                link
                        [attrName=attrValue]
                    子选择器
                        selector > selector

                    子孙选择器
                        selector selector

            css属性:
                attrname: attrvalue;

                heigh: xxxx;
                width: "";
                color: "red";

                w3c 协议规范
                w3school

            选择器
            常用属性
                字体颜色
                    color: red/green/gray
                        rgb(xxx,xx,xx) 0-255
                        #00 00 00
                        00-FF (0-255)
                宽度
                    width:
                高度
                    height:
                定位:
                    position: absolute;
                              relative
            盒子模型:
                标签在页面上占用的空间大小

            bootstrap
                前端css框架
                    定义selector样式 class选择器

                1. 使用
                    引入: link
                2. 如何使用
                    <tagName class="classSelector">

                版本选择:
                    bootstrap前端css框架 js功能


                    选择前端模板(css,js,html)

                基本概念
                    layout:
                        布局
                        container
                        container-fluid

                    grid:
                        栅格
                        行/列(12)
                        <div class="row">
                            <div class="col-2 offset-1"></div>
                            <div class="col-3"></div>
                            <div class="col-6"></div>
                        </div>

        bootsrap js css文件放入到beego项目中放在什么目录下?