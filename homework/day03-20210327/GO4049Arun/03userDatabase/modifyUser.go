package main

import "fmt"

func modifyUser(users []map[string]string)  {
	fmt.Printf("请输入要修改的用户ID:")
	var(
		id string
		flag string
		userName string
		address string
		phoneNumber string
	)
	m:=0
	for m<3{
		fmt.Scan(&id)
		for i := 0; i < len(users); i++ {
			if users[i]["id"] == id{
				//golang map取key的时候是无序的,稍后解决该问题,临时指定了
				fmt.Printf("id:%v,name:%v,tel:%v,addr:%v;",users[i]["id"],users[i]["name"],users[i]["tel"],users[i]["addr"])
				fmt.Println("在数据库中找到用户,您确认要修改用户信息吗?(y/yes/Y/YES)")
				fmt.Scan(&flag)
				if "y"==flag||"yes"==flag||"Y"==flag||"YES"==flag{
					fmt.Printf("请输入用户名(1~10非空字符,不能数字开头):")
					for {
						fmt.Scan(&userName)
						if checkName(userName) {
							users[i]["name"] = userName
							break
						}else {
							fmt.Println("名字输入不符合规范,请重新输入!")
						}
					}
					fmt.Printf("请输入联系方式(11位数字):")
					for  {
						fmt.Scan(&phoneNumber)
						if checkPhoneNumber(phoneNumber) {
							users[i]["tel"] = phoneNumber
							break
						} else {
							fmt.Println("号码输入不符合规范,请重新输入!")
						}
					}
					fmt.Printf("请输入地址(2~20非空字符,不能数字开头):")
					for  {
						fmt.Scan(&address)
						if checkAddr(address) {
							users[i]["addr"] = address
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