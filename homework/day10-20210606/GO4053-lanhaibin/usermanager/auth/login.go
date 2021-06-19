package auth

import (
	"net/http"
	"time"
	"usermanager/model"
	"usermanager/session"
)

func Login(w http.ResponseWriter, r *http.Request, user *model.User) {
	var sid string
	c, err := r.Cookie("sid")
	if err == nil {
		sid = c.Value
	}
	s := session.LoadOrNewSession(sid)
	s.Set("uid", user.Id)
	defer session.DumpSession(&s)
	http.SetCookie(w, &http.Cookie{Name: "sid", Value: s.Id, Path: "/"})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("sid")
	session.DeleteSession(c.Value)
	http.SetCookie(w, &http.Cookie{Name: "sid", Value: "", Expires: time.Unix(0, 0)})
}
