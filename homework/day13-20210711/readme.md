1. 用户管理 数据库修改为ORM
2. 资产管理(增删改查)/角色管理(显示/查询)
3. flag 命令
    --web 启动web服务
    --syncdb 同步数据库
    --force
    --verbose
    --init-user 初始化管理员
        name
        password => 从命令行让用户输入
        存储到db

4. 用户授权控制
    RBAC
    给用户添加角色

    用户 => Role => Module