package session

import "github.com/google/uuid"

type Session struct {
	Id     string
	Values map[string]interface{}
}

func NewSession() Session {
	return Session{
		Id:     uuid.New().String(),
		Values: make(map[string]interface{}),
	}
}

func (s Session) Get(k string) (interface{}, bool) {
	v, ok := s.Values[k]
	return v, ok
}

func (s Session) Set(k string, v interface{}) {
	s.Values[k] = v
}

func (s Session) Delete(k string) {
	delete(s.Values, k)
}
