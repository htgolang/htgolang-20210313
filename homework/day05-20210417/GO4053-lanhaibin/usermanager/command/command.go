package command

type Command struct {
	Desc string
	Cmd  func()
}

var Commands []Command

func Register(desc string, cmd func()) {
	c := Command{
		Desc: desc,
		Cmd:  cmd,
	}
	Commands = append(Commands, c)
}
