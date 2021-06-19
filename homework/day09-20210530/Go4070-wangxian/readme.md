   测试：使用postman
   提交数据用json, 响应数据用json
    a. 添加用户接口 POST
        {
			"id":  1,
            "name" : "xxx",
            "tel" : "xxx",
			"addr": "xxx"
        }

        {
            ok: true/false
        }
    b. 查询用户接口  GET
        queryparams : q =""
        {
            ok: true/false,
            data: [

            ]
        }

    c. 按ID获取用户详情API  GET
        queryparams : id=1
        {
            ok: true/false
            data: {
                id:"",
                name: "",
            }
        }
    d. 删除用户接口  DELETE
        queryparams : id = 1
        {
            ok:true/false
        }
    e. 编辑用户接口  PUT
        {
            id : "",
            name: "",
            tel : "",
            addr: ""
        }

        {
            ok:true/false
        }