package command

type FuncCli interface {
	Name() string
	Execute(args ...interface{})
	GetArgs() []interface{}
}

type BaseFuncCli struct {
}
