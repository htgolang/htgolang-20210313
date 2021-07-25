package templates

import (
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
}
