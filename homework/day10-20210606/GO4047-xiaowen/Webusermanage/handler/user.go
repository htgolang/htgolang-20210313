package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"webusermanage/auth"
	"webusermanage/config"
	"webusermanage/persistence"
	"webusermanage/user"
	"webusermanage/web"
)

func AddHande(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登入
	
	session := web.LoadSession(w, r)
	if _, err := strconv.Atoi(session.Get("uid")); err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}
	if r.Method == http.MethodPost{
		u := config.UserStructure{}
		u.Id = r.PostFormValue("id")
		u.Email = r.PostFormValue("email")
		u.UserName = r.PostFormValue("username")
		u.Password = r.PostFormValue("password")
		u.Phone = r.PostFormValue("phone")
		_, err := user.Add(u)
		if err != nil {
			fmt.Println("add err", err)
			log.Fatalln(err)
		}
		http.Redirect(w,r,"/user/", http.StatusFound)
		return
	}

	tpl, tplerr := template.ParseFiles([]string{"view/add/add.html"}...)
	if tplerr != nil {
		log.Fatalln(tplerr)
	}
	tpl.ExecuteTemplate(w, "add.html", nil)

}

func QueryHande(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登入
	session := web.LoadSession(w, r)
	if _, err := strconv.Atoi(session.Get("uid")); err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}
	r.ParseForm()
	name := r.FormValue("name")
	fmt.Println(name)
	data := user.Query(name)
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
	if _, err := strconv.Atoi(session.Get("uid")); err != nil {
		http.Redirect(w, r, "/login/", http.StatusFound)
		return
	}

	if r.Method == http.MethodPost{
		u := config.UserStructure{}
		u.Id = r.PostFormValue("id")
		u.Email = r.PostFormValue("email")
		u.UserName = r.PostFormValue("username")
		u.Password = r.PostFormValue("password")
		u.Phone = r.PostFormValue("phone")
		r.ParseForm()
		id := r.FormValue("id")
		user.Modif(&u, id)
		http.Redirect(w,r,"/user/", http.StatusFound)
		return
	}
	users := make(map[string]config.UserStructure)
	r.ParseForm()
	id := r.FormValue("id")
	p := persistence.Persistence{StroagePath: config.FilePath}
	err := p.Decode(&users)
	if err != nil{
		log.Fatalln(err)
	}
	for k, v := range users{
		if k == id {
			tpl, tplerr := template.ParseFiles([]string{"view/modify/modify.html"}...)
			if tplerr != nil {
				log.Fatalln(tplerr)
			}
			err = tpl.ExecuteTemplate(w, "modify.html", v)
			if err != nil{
				log.Fatalln(err)
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

	users := user.QueryAll()
	tpl, err := template.ParseFiles([]string{"view/user/user.html"}...)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(w, "user.html", users)
	if err != nil {
		log.Fatalln(err)
	}
}
