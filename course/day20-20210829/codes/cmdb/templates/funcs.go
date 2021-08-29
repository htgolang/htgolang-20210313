package templates

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.AddFuncMap("dateformat", func(t *time.Time) string {
		if t == nil {
			return ""
		}
		return t.Format("2006-01-02")
	})

	web.AddFuncMap("in_string_slice", func(e string, list []string) bool {
		fmt.Println("in string", e, list)
		if list == nil {
			return false
		}
		for _, v := range list {
			if v == e {
				return true
			}
		}
		return false
	})
}
