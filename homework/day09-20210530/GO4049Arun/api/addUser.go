package api

import (
	"fmt"
	"mgr/conf"
	"mgr/model/log"
	"mgr/model/user"
	"net/http"
	"regexp"
)

func AddUserApi(res *ResponseJson){
	fmt.Println("xxx11")
	http.HandleFunc("/adduser/", func(w http.ResponseWriter, r *http.Request) {
		// queryparams -> r.Form
		fmt.Println("xxx22")
		r.ParseForm()
		fmt.Println(r.Method)
		fmt.Println(r.PostForm)
		fmt.Println(r.Form)
		if len(r.PostForm) >0{
			name,flagName := input(&w,res.Data, r.PostFormValue("name"), "name")
			addr,flagAddr := input(&w,res.Data, r.PostFormValue("addr"), "addr")
			phoneNum,flagPhoneNum := input(&w,res.Data, r.PostFormValue("tel"), "phoneNum")
			pwd,flagPwd:=input(&w,res.Data, r.PostFormValue("pwd"),"pwd")
			if flagName&&flagAddr&&flagPhoneNum&&flagPwd{
				res.Data = append(res.Data, &user.User{
					Id: res.Data[len(res.Data)-1].Id +1,
					Name: name,
					Addr: addr,
					Tel: phoneNum,
					Pwd: pwd,
				},
				)
				//ctlUser.Create("json").Write(res.Data)
			}else {
				log.LogRecord("用户接口添加用户失败,请检查原因!")
			}
		}
	})
}

func usersPrint(users []*user.User) {
	for i := 0; i < len(users); i++ {
		//golang map取key的时候是无序的,稍后解决该问题,临时指定了
		//go fmt.Printf中文对齐需要自己单独处理,稍后解决
		fmt.Printf("|id:%-2d|name:%-10s|tel:%-11s|addr:%-20s |\n",users[i].Id,users[i].Name,users[i].Tel,users[i].Addr)
	}
}

func checkExistUser(users []*user.User, name string) bool {
	for i := 0; i < len(users); i++ {
		if name == users[i].Name {
			return true
		}
	}
	return false
}

func checkName(alphabet string) bool {
	//一<\u4e00>~龥<\u9fa5>(基本汉字范围)
	if ok, _ := regexp.MatchString("^[a-zA-Z\u4e00-\u9fa5]\\S{1,9}$", alphabet); !ok {
		return false
	}
	return true
}

func checkAddr(alphabet string) bool {
	//一<\u4e00>~龥<\u9fa5>(基本汉字范围)
	//if ok, _ := regexp.MatchString("^[a-zA-Z\u4e00-\u9fa5]{2,20}$", alphabet); !ok {
	if ok, _ := regexp.MatchString("[\\S]{2,20}$", alphabet); !ok {
		return false
	}
	return true
}

func checkPhoneNumber(phone string) bool {
	if ok, _ := regexp.MatchString("^[0-9]{11}$", phone); !ok {
		return false
	}
	return true
}

func checkPwd(pwd string) bool {
	//特殊字符,大小写字母,数字四选三组合,至少八位
	//正则表达式  ^(?![a-zA-Z]+$)(?![A-Z0-9]+$)(?![A-Z\\W_]+$)(?![a-z0-9]+$)(?![a-z\\W_]+$)(?![0-9\\W_]+$)[a-zA-Z0-9\\W_]{8,}$
	if ok, _ := regexp.MatchString(fmt.Sprintf("^[0-9]{%d}$",conf.Info.PwdLen), pwd); !ok {
		return false
	}
	return true
}

func input(w *http.ResponseWriter,users []*user.User, value string, flag string) (string,bool) {
	var text string
	if "name" == flag{
		if checkName(text) {
			// 验证 用户名不能重复
			if checkExistUser(users, text) {
				MyMarshalError(w,fmt.Sprintf("用户名%s已存在,请重新提交!\n", text) )
				return value,false
			}
			return value,true
		} else {
			fmt.Println("名字输入不符合规范,请重新输入!")
		}
	} else if "phoneNum" == flag {
		if checkPhoneNumber(text) {
			return value,true
		} else {
			fmt.Println("号码输入不符合规范,请重新输入!")
		}
	}else if "pwd" == flag {
		if checkPwd(text) {
			return value,true
		} else {
			fmt.Printf("密码输入不符合%d个密码长度要求,请重新输入!\n",conf.Info.PwdLen)
		}
	} else {
		if checkAddr(text) {
			return value,true
		} else {
			fmt.Println("地址输入不符合规范,请重新输入!")
		}
	}
	return text,true
}
