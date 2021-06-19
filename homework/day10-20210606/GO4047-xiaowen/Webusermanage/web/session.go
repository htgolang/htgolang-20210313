package web

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"webusermanage/utils"

	"github.com/google/uuid"
)

const (
	sessionName = "sid"
	sessionDir  = "temp/sessions"
)

type Session struct {
	ID     string
	Values map[string]string
}

func (s *Session) Get(name string) string {
	return s.Values[name]
}

func (s *Session) Set(name, value string) {
	s.Values[name] = value
}

func LoadSession(w http.ResponseWriter, r *http.Request) *Session {
	if cookie, err := r.Cookie(sessionName); err == nil {
		if cookie != nil && cookie.Value != "" && utils.IsLetterOrNumber(cookie.Value) {
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
		Values: make(map[string]string),
	}
}

func DumpSession(w http.ResponseWriter, r *http.Request, session *Session) error {
	sessionPath := path.Join(sessionDir, session.ID)
	file, err := os.Create(sessionPath)
	if err == nil {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoderErr := encoder.Encode(session)
		return encoderErr
	} else {
		return err
	}
}

func DeleteSession(w http.ResponseWriter, r *http.Request, session *Session) error {
	http.SetCookie(w, &http.Cookie{
		Name:   sessionName,
		Path:   "/",
		Value:  "",
		MaxAge: -1,
	})
	sessionPath := path.Join(sessionDir, session.ID)
	return os.Remove(sessionPath)
}
