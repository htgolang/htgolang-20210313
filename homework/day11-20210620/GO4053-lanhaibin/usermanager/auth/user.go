package auth

import (
	"net/http"
	"usermanager/model"
	"usermanager/service"
	"usermanager/session"
)

func GetCurrentUser(r *http.Request) model.User {
	c, _ := r.Cookie("sid")
	sid := c.Value
	s, _ := session.LoadSession(sid)
	id, _ := s.Get("uid")
	user, _ := service.Us.GetUserById(int(id.(float64)))
	return user
}
