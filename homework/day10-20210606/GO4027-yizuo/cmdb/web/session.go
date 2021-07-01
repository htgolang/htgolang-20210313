package web

import (
	"cmdb/utils"
	"encoding/json"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
)

const (
	sessionName = "sid"
	sessionDir  = "temp/sessions"
)

type Session struct {
	ID     string
	Values map[string]interface{}
}

func (s Session) Get(name string) interface{} {
	return s.Values[name]
}

func (s Session) Set(name string, value interface{}) {
	s.Values[name] = value
}

func LoadSession(w http.ResponseWriter, r *http.Request) *Session {
	if cookie, err := r.Cookie(sessionName); err == nil {
		if cookie.Value != "" && utils.IsLetterOrNumer(cookie.Value) {
			sessionPath := path.Join(sessionDir, cookie.Value)
			if file, err := os.Open(sessionPath); err == nil {
				defer file.Close()
				var session Session
				decoder := json.NewDecoder(file)
				if err := decoder.Decode(&session); err == nil {
					return &session
				}
			}
		}
	}

	sid := strings.ReplaceAll(uuid.New().String(), "-", "")
	http.SetCookie(w, &http.Cookie{
		Name:  sessionName,
		Path:  "/",
		Value: sid,
	})

	return &Session{
		ID:     sid,
		Values: make(map[string]interface{}),
	}
}

func DumpSession(w http.ResponseWriter, r *http.Request, session *Session) error {
	fpath := path.Join(sessionDir, session.ID)
	if file, err := os.Create(fpath); err != nil {
		return err
	} else {
		defer file.Close()
		encoder := json.NewEncoder(file)
		err := encoder.Encode(session)
		return err
	}

}

func DeleteSession(w http.ResponseWriter, r *http.Request, session *Session) error {
	http.SetCookie(w, &http.Cookie{
		Name:   sessionName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	fpath := path.Join(sessionDir, session.ID)
	return os.Remove(fpath)

}
