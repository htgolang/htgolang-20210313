package web

import (
	"bytes"
	"cmdb/utils"
	"encoding/gob"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

const (
	tokenName = "token"
	signName  = "sign"
)

type Token struct {
	ID     string
	Values map[string]interface{}
}

func (t *Token) Set(name string, value interface{}) {
	t.Values[name] = value
}

func (t *Token) Get(name string) interface{} {
	return t.Values[name]
}

func LoadToken(w http.ResponseWriter, r *http.Request) *Token {
	tokenValue, err := r.Cookie(tokenName)
	if err == nil && tokenValue != nil && tokenValue.Value != "" {
		signValue, err := r.Cookie(signName)
		if err == nil && signValue != nil && signValue.Value != "" {
			// signValue
			// tokenValue 计算签名
			token, err := utils.Base64Decode(tokenValue.Value)
			if err == nil {
				sign := utils.Signature(token)
				if sign == signValue.Value {
					var t Token
					buffer := bytes.NewBuffer(token)
					decoder := gob.NewDecoder(buffer)
					if err := decoder.Decode(&t); err == nil {
						return &t
					}
				}
			}
		}
	}
	return &Token{
		ID:     strings.ReplaceAll(uuid.New().String(), "-", ""),
		Values: make(map[string]interface{}),
	}
}

func DumpToken(w http.ResponseWriter, r *http.Request, token *Token) error {
	// Token = gob编码 => 签名
	// token base64编码
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	if err := encoder.Encode(token); err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  tokenName,
		Path:  "/",
		Value: utils.Base64Encode(buffer.Bytes()),
	})

	http.SetCookie(w, &http.Cookie{
		Name:  signName,
		Path:  "/",
		Value: utils.Signature(buffer.Bytes()),
	})
	return nil
}
