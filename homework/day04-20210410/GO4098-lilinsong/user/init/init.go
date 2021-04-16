package init

import (
	mt "user/commands"
	co "user/consoles"
)

func init() {
	co.Register(mt.Add)
	co.Register(mt.Del)
	co.Register(mt.Query)
	co.Register(mt.Modify)

}
