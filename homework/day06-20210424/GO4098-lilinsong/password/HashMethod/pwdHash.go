package HashMethod

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type pwdHash struct {
	SaltLength     int
	passwordLength int
	password       string
	chars          string
	salt           string
	hash           *HashEncryption
}

type pwdMethod interface {
	randSalt()
	CheckPassword(pwd, hash, t string, pg *sync.WaitGroup) bool
	HashPassword(pwd, hash string, slat int, pg *sync.WaitGroup) (string, error)
}

type PwdHashMethod interface {
	pwdMethod() pwdMethod
}

func (h *HashEncryption) pwdMethod() pwdMethod {
	return &pwdHash{
		hash: h,
	}
}

func NewPwdHash() *pwdHash {
	return &pwdHash{
		password:       "",
		SaltLength:     6,
		passwordLength: 16,
		chars:          "123456789asdfghjklpoiuytrewqzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM",
		hash: &HashEncryption{
			startCha: make(chan struct{}, 1),
			stopCha:  make(chan struct{}, 1),
		},
	}
}

func (p *pwdHash) randSalt() {
	for i := 0; i < p.SaltLength; i++ {
		p.salt += string(p.chars[rand.Intn(len(p.chars))])
	}
}
func (p *pwdHash) CheckPassword(pwd, hash, t string, pg *sync.WaitGroup) bool {
	defer pg.Done()
	pos := strings.LastIndex(hash, "=")
	if pos < 0 {
		return false
	}

	p.salt = hash[pos:]
	p.hash.wg.Add(1)
	go func() {

		p.password = p.hash.method(t, "", []byte(pwd))
	}()

	p.hash.startCha <- struct{}{}
	p.hash.stopCha <- struct{}{}

	p.hash.wg.Wait()
	return fmt.Sprintf("%s%s", p.password, p.salt) == hash
}
func (p *pwdHash) HashPassword(pwd, hash string, slat int, pg *sync.WaitGroup) (string, error) {
	defer pg.Done()
	p.SaltLength = slat
	p.randSalt()
	if len(p.password) <= p.passwordLength {
		p.hash.wg.Add(1)
		go func() {
			p.password = p.hash.method(hash, "", []byte(pwd))
		}()

		p.hash.startCha <- struct{}{}
		p.hash.stopCha <- struct{}{}
	}
	p.hash.wg.Wait()

	return fmt.Sprintf("%s=%s", p.password, p.salt), nil
}
