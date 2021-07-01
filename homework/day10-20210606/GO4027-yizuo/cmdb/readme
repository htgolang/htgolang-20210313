1. 用户管理改成web
   测试：使用postman
   提交数据用json, 响应数据用json
    a. 添加用户接口
        {
            "name" : "xxx",
            "tel" : "xxx",
        }

        {
            ok: true/false
        }
    b. 查询用户接口
        queryparams : q =""
        {
            ok: true/false,
            data: [

            ]
        }

    c. 按ID获取用户详情API
        queryparams : id=1
        {
            ok: true/false
            data: {
                id:"",
                name: "",
            }
        }
    d. 删除用户接口
        queryparams : id = 1
        {
            ok:true/false
        }
    e. 编辑用户接口
        {
            id : "",
            name: "",
            tel : "",
            ...
        }

        {
            ok:true/false
        }