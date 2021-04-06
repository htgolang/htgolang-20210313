yizuodeMacBook-Pro:3_用户管理 yizuo$ go run main.go 
请输入想要执行的操作:(add/del/mod/query/list/quit) 
add
请输入用户名称：
leimu
请输入用户地址：
shanghai
请输入联系方式：
123456
[map[addr:Wuhan id:1 name:yizuo tel:66666666] map[addr:shanghai id:2 name:leimu tel:123456]]
请输入想要执行的操作:(add/del/mod/query/list/quit) 
list
[map[addr:Wuhan id:1 name:yizuo tel:66666666] map[addr:shanghai id:2 name:leimu tel:123456]]
请输入想要执行的操作:(add/del/mod/query/list/quit) 
query
请输入要查询的信息:(默认扫描所有数据检索) 
shanghai
查询到的用户数据为：
map[addr:shanghai id:2 name:leimu tel:123456]
请输入想要执行的操作:(add/del/mod/query/list/quit) 
query
请输入要查询的信息:(默认扫描所有数据检索) 
yizuo
查询到的用户数据为：
map[addr:Wuhan id:1 name:yizuo tel:66666666]
请输入想要执行的操作:(add/del/mod/query/list/quit) 
del
请输入要删除的用户ID: 
2
你要删除的的用户数据为：map[addr:shanghai id:2 name:leimu tel:123456]
确认要删除用户吗？：(yes/no ; default no)
yes
[map[addr:Wuhan id:1 name:yizuo tel:66666666]]
请输入想要执行的操作:(add/del/mod/query/list/quit) 
quit