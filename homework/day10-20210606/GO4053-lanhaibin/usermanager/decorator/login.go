package decorator

import (
	"net/http"
	"usermanager/session"
)

func LoginRequired(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("sid")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		s, err := session.LoadSession(c.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		if _, ok := s.Get("uid"); !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		f(w, r)
	}
}
