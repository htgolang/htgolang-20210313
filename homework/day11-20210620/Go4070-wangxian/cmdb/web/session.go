package web

import (
	"cmdb/utils"
	"encoding/gob"
	"net/http"
	"os"
	"path"

	"github.com/google/uuid"
)

var (
	sessinKey  string = "sid"
	sessionDir string = "tmp/session/"
)

type Session struct {
	Id     string
	Values map[string]interface{}
}

func (s Session) Get(name string) interface{} {
	return s.Values[name]
}

func (s *Session) Set(name string, value interface{}) {
	s.Values[name] = value
}

//从cookie中读取sid，根据sid加载session文件
func LoadSession(w http.ResponseWriter, r *http.Request) *Session {
	//先获取cookie
	cookie, err := r.Cookie(sessinKey)
	if err == nil && cookie.Value != "" && utils.IsLetterAndNum(cookie.Value) { //判断sid存在，且值不为空,且值符合规范
		fpath := path.Join(sessionDir, cookie.Value)
		if sfile, err := os.Open(fpath); err == nil { //打开session文件
			defer sfile.Close()
			var session Session
			//对session文件解码
			decoder := gob.NewDecoder(sfile)
			if err := decoder.Decode(&session); err == nil {
				//session有效
				return &session
			}
		}
	}
	//sid不存在，第一次访问
	//在响应中设置cookie sid
	sid := uuid.New().String()
	http.SetCookie(w, &http.Cookie{
		Name:  sessinKey,
		Value: sid,
		Path:  "/",
	})

	return &Session{Id: sid, Values: map[string]interface{}{}}
}

//把session保存到文件中
func DumpSession(w http.ResponseWriter, r *http.Request, session *Session) error {
	fpath := path.Join(sessionDir, session.Id)
	sfile, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer sfile.Close()
	encoder := gob.NewEncoder(sfile)
	err = encoder.Encode(session)
	return err
}

//删除session
func DeleteSession(w http.ResponseWriter, r *http.Request, session *Session) error {
	http.SetCookie(w, &http.Cookie{
		Name:   sessinKey,
		Value:  "",
		MaxAge: -1,
	})

	//删除session文件
	fpath := path.Join(sessionDir, session.Id)
	return os.Remove(fpath)
}
