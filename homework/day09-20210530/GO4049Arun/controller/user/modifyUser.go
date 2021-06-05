package user

import (
	"fmt"
	"mgr/conf"
	"mgr/model/user"
)

func ModifyUser(users []*user.User)  {
	fmt.Printf("请输入要修改的用户ID:")
	var(
		flag string
		user user.User
	)
	m:=0
	//users = Create("gob").Read()
	users = Create("json").Read()
	//users = Create("csv").Read()
	for m<3{
		fmt.Scan(&user.Id)
		for i := 0; i < len(users); i++ {
			if users[i].Id == user.Id{
				//golang map取key的时候是无序的,稍后解决该问题,临时指定了
				fmt.Printf("id:%v,name:%v,tel:%v,addr:%v;",users[i].Id,users[i].Name,users[i].Tel,users[i].Addr)
				fmt.Println("在数据库中找到用户,您确认要修改用户信息吗?(y/yes/Y/YES)")
				fmt.Scan(&flag)
				if "y"==flag||"yes"==flag||"Y"==flag||"YES"==flag{
					if CurrentUser == users[i].Name{
						fmt.Printf("您是当前用户:%s,请先输入要修改的新密码:",CurrentUser)
						for  {
							fmt.Scan(&user.Pwd)
							if checkPwd(user.Pwd) {
								users[i].Pwd = user.Pwd
								break
							} else {
								fmt.Printf("密码输入不符合%d个密码长度要求,请重新输入!\n",conf.Info.PwdLen)
							}
						}
					}
					fmt.Printf("请输入用户名(1~10非空字符,不能数字开头):")
					for {
						fmt.Scan(&user.Name)
						if checkName(user.Name) {
							users[i].Name = user.Name
							break
						}else {
							fmt.Println("名字输入不符合规范,请重新输入!")
						}
					}
					fmt.Printf("请输入联系方式(11位数字):")
					for  {
						fmt.Scan(&user.Tel)
						if checkPhoneNumber(user.Tel) {
							users[i].Tel = user.Tel
							break
						} else {
							fmt.Println("号码输入不符合规范,请重新输入!")
						}
					}
					fmt.Printf("请输入地址(2~20非空字符,不能数字开头):")
					for  {
						fmt.Scan(&user.Addr)
						if checkAddr(user.Addr) {
							users[i].Addr = user.Addr
							break
						} else {
							fmt.Println("地址输入不符合规范,请重新输入!")
						}
					}

				}
				goto done
			}
		}
		if m==2{
			fmt.Printf("您次数已超过上限%v,即将退出程序!",m+1)
			return
		}
		m++
		fmt.Println("您输入的用户ID不存在,请重新输入!")
		continue
	}
done:
	usersPrint(users)
}