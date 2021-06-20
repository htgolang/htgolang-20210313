package pkg

import (
	_ "embed"
)

// 找对应的embed文件 => 相对当前go文件所在目录, 并且不能设置其父目录

//go:embed pkg.name
var Name string

//go:embed params/pkg.version
var Version string
