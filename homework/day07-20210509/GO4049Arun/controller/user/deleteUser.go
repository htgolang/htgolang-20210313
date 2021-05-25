package user

import (
	"fmt"
	"mgr/model/user"
)

func DeleteUser(users []*user.User)  {
	fmt.Printf("请输入要删除的用户ID:")
	var(
		id int
		flag string
	)
	m:=0
	//users = Create("gob").Read()
	users = Create("json").Read()
	//users = Create("csv").Read()
	for m<3{
		fmt.Scan(&id)
		for i := 0; i < len(users); i++ {
			if users[i].Id == id{
				//golang map取key的时候是无序的,稍后解决该问题,临时指定了
				fmt.Printf("id:%v,name:%v,addr:%v,tel:%v;",users[i].Id,users[i].Name,users[i].Addr,users[i].Tel)
				fmt.Println("您确认要删除吗?(y/yes/Y/YES)")
				fmt.Scan(&flag)
				if "y"==flag||"yes"==flag||"Y"==flag||"YES"==flag{
					if 1!=len(users){
						users = append(users[:i],users[i+1:]...)
					}else {
						users = users[:0]
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
	//Create("gob").Write(users)
	Create("json").Write(users)
	//Create("csv").Write(users)
	usersPrint(users)
}