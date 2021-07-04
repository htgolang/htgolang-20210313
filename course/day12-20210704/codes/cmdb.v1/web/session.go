package web

import (
	"cmdb/utils"
	"encoding/gob"
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

func (s *Session) Set(name string, value interface{}) {
	s.Values[name] = value
}

func (s *Session) Get(name string) interface{} {
	return s.Values[name]
}

func LoadSession(w http.ResponseWriter, r *http.Request) *Session { // 从cookie读取sid 加载文件
	if cookie, err := r.Cookie(sessionName); err == nil {
		if cookie != nil && cookie.Value != "" && utils.IsLetterOrNumer(cookie.Value) {
			fpath := path.Join(sessionDir, cookie.Value)
			if fhandler, err := os.Open(fpath); err == nil {
				defer fhandler.Close()
				var session Session
				decoder := gob.NewDecoder(fhandler)
				if err := decoder.Decode(&session); err == nil {
					// session有效
					return &session
				}
				// 删除文件
			}
		}
	}
	sid := strings.ReplaceAll(uuid.New().String(), "-", "")
	http.SetCookie(w, &http.Cookie{
		Name:  sessionName,
		Value: sid,
		Path:  "/",
	})

	// 生成Session
	return &Session{
		ID:     sid,
		Values: make(map[string]interface{}),
	}

}

func DumpSession(w http.ResponseWriter, r *http.Request, session *Session) error { // 将session存储到对应的文中
	fpath := path.Join(sessionDir, session.ID)
	if fhandler, err := os.Create(fpath); err != nil {
		return err
	} else {
		defer fhandler.Close()
		encoder := gob.NewEncoder(fhandler)
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
