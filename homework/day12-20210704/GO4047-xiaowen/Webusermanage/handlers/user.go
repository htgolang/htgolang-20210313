package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"webusermanage/models"
	"webusermanage/services"
	"webusermanage/utils"

	"github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

func AddHande(cc *web.Controller) {
	var u models.User
	u.Email = cc.GetString("email")
	u.Name = cc.GetString("username")
	hashpwd, err := bcrypt.GenerateFromPassword([]byte(cc.GetString("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	u.Password = string(hashpwd)
	u.Phone = cc.GetString("phone")
	role := cc.GetString("role")
	if role == "Admin" {
		u.RoleId = 1
	} else if role == "Operation" {
		u.RoleId = 2
	}
	fmt.Printf("%#v\n", &u)
	_, err = services.Add(&u)
	if err != nil {
		log.Fatalln(err)
	}
}

//这个方法未用
func QueryHande() []models.User {
	users := services.QueryAll()
	return users
}

func ModifyHande(id int, c *web.Controller) {
	var i, role_id int
	var name, phone, email string
	qsql := "select id,name,phone,email,role_id from user where id=?"
	dsn, _ := web.AppConfig.String("mysql::dsn")
	db, oerr := sql.Open("mysql", dsn)
	if oerr != nil {
		log.Println(oerr)
	}
	defer db.Close()
	row, qerr := db.Query(qsql, id)
	if qerr != nil {
		log.Println(qerr)
	}
	defer row.Close()
	sourceerr := ""
	newerr := ""
	//post方法走这
	if c.Ctx.Input.IsPost() {
		u := models.User{}
		u.Id = id
		u.Email = c.GetString("email")
		u.Name = c.GetString("username")
		flags := utils.SourcePwdValidate(c.GetString("sourcePassword"), c.GetString("username"))
		//原密码错误时处理
		if !flags {
			sourceerr = "密码错误"
			for row.Next() {
				if rerr := row.Scan(&i, &name, &phone, &email, &role_id); rerr == nil {
					if role_id == 1 {
						role_name := "Admin"
						c.TplName = "modify/modify.html"
						c.Data["Users"] = models.ModifyTpl{i, name, phone, email, role_name, sourceerr, newerr}
					} else if role_id == 2 {
						role_name := "Operation"
						c.TplName = "modify/modify.html"
						c.Data["Users"] = models.ModifyTpl{i, name, phone, email, role_name, sourceerr, newerr}
					}
				}
			}
			return
		}
		u.Phone = c.GetString("phone")
		newPassword, _ := utils.Genpwd(c.GetString("newPassword"))
		u.Password, flags = utils.Comparepwd([]byte(newPassword), c.GetString("validatePassword"))
		//修改的密码不一致时处理
		if !flags {
			newerr = "密码不一致"
			for row.Next() {
				if rerr := row.Scan(&i, &name, &phone, &email, &role_id); rerr == nil {
					if role_id == 1 {
						role_name := "Admin"
						c.TplName = "modify/modify.html"
						c.Data["Users"] = models.ModifyTpl{i, name, phone, email, role_name, sourceerr, newerr}
					} else if role_id == 2 {
						role_name := "Operation"
						c.TplName = "modify/modify.html"
						c.Data["Users"] = models.ModifyTpl{i, name, phone, email, role_name, sourceerr, newerr}
					}
				}
			}
			return
		} else if flags {
			id := c.GetString("id")
			services.Modif(&u, id)
			c.Redirect(web.URLFor("UserController.Query"), http.StatusFound)
			return
		}
	}

	//显示修改页面
	for row.Next() {
		if rerr := row.Scan(&i, &name, &phone, &email, &role_id); rerr == nil {
			if role_id == 1 {
				role_name := "Admin"
				c.TplName = "modify/modify.html"
				c.Data["Users"] = models.ModifyTpl{i, name, phone, email, role_name, sourceerr, newerr}
			} else if role_id == 2 {
				role_name := "Operation"
				c.TplName = "modify/modify.html"
				c.Data["Users"] = models.ModifyTpl{i, name, phone, email, role_name, sourceerr, newerr}
			}
		}
	}
}

func RoleHande() []models.Roles {
	var id int
	var role_name string
	var temp []models.Roles
	qsql := "select DISTINCT user.role_id,role.name from user, role WHERE user.role_id = role.id;"
	dsn, _ := web.AppConfig.String("mysql::dsn")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	row, err := db.Query(qsql)
	if err != nil {
		log.Fatalln(err)
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&id, &role_name)
		role := models.Roles{id, role_name}
		temp = append(temp, role)
	}
	return temp
}
