package handler

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"webusermanage/auth"
	"webusermanage/config"
	"webusermanage/user"
	"webusermanage/utils"
	"webusermanage/web"
)

func AddHande(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登入
	session := web.LoadSession(w, r)
	if _, err := strconv.Atoi(session.Get("uid")); err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}
	//post方法走着
	if r.Method == http.MethodPost{
		u := config.UserStructure{}
		u.Id = r.PostFormValue("id")
		u.Email = r.PostFormValue("email")
		u.UserName = r.PostFormValue("username")
		u.Password = r.PostFormValue("password")
		u.Phone = r.PostFormValue("phone")
		role := r.PostFormValue("role")
		if role == "Admin" {
			u.Role_id = 1
		} else if role == "Operation"{
			u.Role_id = 2
		}
		_, err := user.Add(u)
		if err != nil {
			fmt.Println("add err", err)
			log.Fatalln(err)
		}
		http.Redirect(w,r,"/user/", http.StatusFound)
		return
	}

	//显示添加页面
	tpl, tplerr := template.ParseFiles([]string{"view/add/add.html"}...)
	if tplerr != nil {
		log.Fatalln(tplerr)
	}
	tpl.ExecuteTemplate(w, "add.html", nil)

}

//这个方法未用
func QueryHande(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登入
	session := web.LoadSession(w, r)
	if _, err := strconv.Atoi(session.Get("uid")); err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}
	r.ParseForm()
	id := r.FormValue("id")
	data := user.Query(id)
	fmt.Fprint(w, data)
}

func DelHande(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登入
	session := web.LoadSession(w, r)
	if _, err := strconv.Atoi(session.Get("uid")); err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}
	r.ParseForm()
	id := r.FormValue("id")
	user.Delete(id)
	http.Redirect(w,r,"/user/",http.StatusFound)
	
}

func ModifyHande(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登入
	session := web.LoadSession(w, r)
	if _, aterr := strconv.Atoi(session.Get("uid")); aterr != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	r.ParseForm()
	id, aerr := strconv.Atoi(r.FormValue("id"))
	if aerr != nil{
		log.Println(aerr)
	}
	var i,role_id int
	var name,phone,email string
	qsql := "select id,name,phone,email,role_id from user where id=?"
	db, oerr := sql.Open("mysql",config.DSN)
	if oerr != nil{
		log.Println(oerr)
	}
	defer db.Close()
	row, qerr := db.Query(qsql,id)
	if qerr != nil{
		log.Println(qerr)
	}
	defer row.Close()
	sourceerr := ""
	newerr := ""
	//post方法走这
	if r.Method == http.MethodPost{
		u := config.User{}
		u.Id, _ = strconv.Atoi(r.PostFormValue("id"))
		u.Email = r.PostFormValue("email")
		u.Name = r.PostFormValue("username")
		flags := utils.SourcePwdValidate(r.PostFormValue("sourcePassword"), r.PostFormValue("username")) 
		//原密码错误时处理
		if !flags{
			sourceerr = "密码错误"
			for row.Next(){
				if rerr := row.Scan(&i,&name,&phone,&email,&role_id); rerr == nil{
					if role_id == 1{
						role_name := "Admin"
						tpl, tplerr := template.ParseFiles([]string{"view/modify/modify.html"}...)
						if tplerr != nil {
							log.Fatalln(tplerr)
						}
						terr := tpl.ExecuteTemplate(w, "modify.html", struct{
							Id int
							Name string
							Phone string
							Email string
							Role_name string
							SourceErr string
							NewErr string
						}{Id: i,
						Name: name,
						Phone: phone,
						Email: email,
						Role_name: role_name,
						SourceErr: sourceerr,
						NewErr: newerr,})
						if terr != nil{
							log.Fatalln(terr)
						}
					} else if role_id == 2{
						role_name := "Operation"
						tpl, tplerr := template.ParseFiles([]string{"view/modify/modify.html"}...)
						if tplerr != nil {
							log.Fatalln(tplerr)
						}
						terr := tpl.ExecuteTemplate(w, "modify.html", struct{
							Id int
							Name string
							Phone string
							Email string
							Role_name string
							SourceErr string
							NewErr string
						}{Id: i,
						Name: name,
						Phone: phone,
						Email: email,
						Role_name: role_name,
						SourceErr: sourceerr,
						NewErr: newerr,})
						if terr != nil{
							log.Fatalln(terr)
						}
					}
				}
			}
			return
		}
		u.Phone = r.PostFormValue("phone")
		newPassword, _ :=  utils.Genpwd(r.PostFormValue("newPassword")) 
		u.Password, flags =  utils.Comparepwd([]byte(newPassword), r.PostFormValue("validatePassword"))
		fmt.Println("source flags:", flags)
		//修改的密码不一致时处理
		if !flags{
			sourceerr = ""
			newerr = "密码不一致"
			for row.Next(){
				if rerr := row.Scan(&i,&name,&phone,&email,&role_id); rerr == nil{
					if role_id == 1{
						role_name := "Admin"
						tpl, tplerr := template.ParseFiles([]string{"view/modify/modify.html"}...)
						if tplerr != nil {
							log.Fatalln(tplerr)
						}
						terr := tpl.ExecuteTemplate(w, "modify.html", struct{
							Id int
							Name string
							Phone string
							Email string
							Role_name string
							SourceErr string
							NewErr string
						}{Id: i,
						Name: name,
						Phone: phone,
						Email: email,
						Role_name: role_name,
						SourceErr: sourceerr,
						NewErr: newerr,})
						if terr != nil{
							log.Fatalln(terr)
						}
					} else if role_id == 2{
						role_name := "Operation"
						tpl, tplerr := template.ParseFiles([]string{"view/modify/modify.html"}...)
						if tplerr != nil {
							log.Fatalln(tplerr)
						}
						terr := tpl.ExecuteTemplate(w, "modify.html", struct{
							Id int
							Name string
							Phone string
							Email string
							Role_name string
							SourceErr string
							NewErr string
						}{Id: i,
						Name: name,
						Phone: phone,
						Email: email,
						Role_name: role_name,
						SourceErr: sourceerr,
						NewErr: newerr,})
						if terr != nil{
							log.Fatalln(terr)
						}
					}
				}
			}
			return
		} else if flags {
			r.ParseForm()
			id := r.FormValue("id")
			user.Modif(&u, id)
			http.Redirect(w,r,"/user/", http.StatusFound)
			return
		}
	}


	//显示修改页面
	for row.Next(){
		if rerr := row.Scan(&i,&name,&phone,&email,&role_id); rerr == nil{
			if role_id == 1{
				role_name := "Admin"
				tpl, tplerr := template.ParseFiles([]string{"view/modify/modify.html"}...)
				if tplerr != nil {
					log.Fatalln(tplerr)
				}
				terr := tpl.ExecuteTemplate(w, "modify.html", struct{
					Id int
					Name string
					Phone string
					Email string
					Role_name string
					SourceErr string
					NewErr string
					
				}{Id: i,
				Name: name,
				Phone: phone,
				Email: email,
				Role_name: role_name,
				SourceErr: sourceerr,
				NewErr: newerr,})
				if terr != nil{
					log.Fatalln(terr)
				}
			} else if role_id == 2{
				role_name := "Operation"
				tpl, tplerr := template.ParseFiles([]string{"view/modify/modify.html"}...)
				if tplerr != nil {
					log.Fatalln(tplerr)
				}
				terr := tpl.ExecuteTemplate(w, "modify.html", struct{
					Id int
					Name string
					Phone string
					Email string
					Role_name string
					SourceErr string
					NewErr string
				}{Id: i,
				Name: name,
				Phone: phone,
				Email: email,
				Role_name: role_name,
				SourceErr: sourceerr,
				NewErr: newerr,})
				if terr != nil{
					log.Fatalln(terr)
				}
			}
		}
	}
}

func DatilHande(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登入
	session := web.LoadSession(w, r)
	if _, err := strconv.Atoi(session.Get("uid")); err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}
	r.ParseForm()
	id := r.FormValue("id")
	flag, data := user.GetDtail(id)

	fmt.Fprintf(w, "{\"ok\": %v, \"data\": %#v}", flag, data)
}

func LoginHande(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登入
	session := web.LoadSession(w, r)
	if _, err := strconv.Atoi(session.Get("uid")); err == nil {
		http.Redirect(w, r, "/user/", http.StatusFound)
		return
	}

	username := ""
	password := ""
	err := ""
	// 验证登录
	if r.Method == http.MethodPost {
		username = r.PostFormValue("username")
		password = r.PostFormValue("password")
		if id, flag := auth.Auth(username, password); flag {
			session.Set("uid", id)
			web.DumpSession(w, r, session)
			http.Redirect(w, r, "/user/", http.StatusFound)
			return
		}
		err = "用户名或密码错误"
	}
	tpl, tplerr := template.ParseFiles([]string{"view/login/login.html"}...)
	if tplerr != nil {
		log.Fatalln(tplerr)
	}
	tpl.ExecuteTemplate(w, "login.html", struct {
		name string
		Err  string
	}{username, err})
}

func LogoutHande(w http.ResponseWriter, r *http.Request) {
	web.DeleteSession(w, r, web.LoadSession(w, r))
	http.Redirect(w, r, "/login/", http.StatusFound)
}

func UserHande(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登入
	session := web.LoadSession(w, r)
	if _, err := strconv.Atoi(session.Get("uid")); err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}
	uid := session.Get("uid") 

	users := user.QueryAll()
	currentUser := utils.GetCurrentUser(uid)

	tpl, err := template.ParseFiles([]string{"view/user/user.html", "view/header/header.html"}...)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "user.html", struct{
		CurrentUser *config.UserStructure
		Users []config.UserStructure
	}{currentUser, users})
	if err != nil {
		log.Fatalln(err)
	}
}

func RoleHande(w http.ResponseWriter, r *http.Request){
	//判断是否已经登入
	session := web.LoadSession(w, r)
	if _, err := strconv.Atoi(session.Get("uid")); err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}
	var id int
	var role_name string
	type Roles struct{
		Role_id int
		Role_name string
	}
	var temp []Roles
	qsql := "select DISTINCT user.role_id,role.name from user, role WHERE user.role_id = role.id;"
	db, err := sql.Open("mysql",config.DSN)
	if err != nil{
		log.Println(err)
	}
	defer db.Close()
	row, err := db.Query(qsql)
	if err != nil{
		log.Println(err)
	}
	defer row.Close()
	for row.Next(){
		row.Scan(&id, &role_name)
		role := Roles{id, role_name}
		temp = append(temp, role)
	}
	
	
	tpl, err := template.ParseFiles([]string{"view/role/role.html"}...)
	if err != nil{
		log.Println(err)
	}
	err = tpl.ExecuteTemplate(w, "role.html", temp)
	if err != nil{
		log.Println(err)
	}
}