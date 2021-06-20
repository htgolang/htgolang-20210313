1. 将用户增/删/改/查 => 改到DB
2. 部门/资产 => 增删改查管理
3. 【选做】headers中添加菜单
   用户管理，部门管理，角色管理，资产管理
   对应用户登陆后，只能根据只能根据对应的角色显示有权限的菜单

4. 【选做】角色 => 只有列表显示/无增删改
   角色表加modules字段 text = json字符串["userManager", "departmentManager", "roleManager, "assertManager"]
    默认设置角色:
        管理员 userManager, departmentManager, roleManager
        运维人员 assetManager权限

    每个模块对应相应的菜单
        userManager => 用户管理菜单
        departmentManager => 部门管理菜单
        roleManager => 角色管理菜单
        assertManager => 资产管理菜单
5. 用户添加 用户密码
    用户修改 => 不会修改密码
    每个人具有修改自己密码的功能(点击登陆后菜单上的 欢迎xxx 跳转到修改密码页面)
    原密码/新密码/确认密码 => 修改数据库中password